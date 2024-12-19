package v1

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/db"
	gen "gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/gen/proto/v1"
)

const (
	gRPCGatewayCookie      = "grpcgateway-cookie"
	authorizationCookieKey = "Authorization"
	userIDMDKey            = "user_id"
	adminID                = "admin"
)

var (
	permissionDeniedErr = fmt.Errorf("permission denied")
)

type Authorizer struct {
	gen.UnimplementedAuthorizerServer

	Repo db.Repository

	JWTSecret  string
	AdminToken string
}

type Claims struct {
	jwt.StandardClaims

	UserID string `json:"user_id"`
}

// endpointPermissionValidators is a list of regexes and corresponding validators
// that check if a user has permission to access an endpoint.
var endpointPermissionValidators = []struct {
	regex     string
	onlyAdmin bool
}{
	{
		regex:     gen.Authorizer_IsAdmin_FullMethodName,
		onlyAdmin: true,
	},
	{
		regex:     `\/fitness_aggregator\.v1\.Authorizer\/*`,
		onlyAdmin: false,
	},
	{
		regex:     `\/fitness_aggregator\.v1\.SearchEngine\/*`,
		onlyAdmin: false,
	},
	{
		regex:     `\/fitness_aggregator\.v1\.FitnessAggregator\/CreateClient`,
		onlyAdmin: false,
	},
	{
		regex:     `\/fitness_aggregator\.v1\.FitnessAggregator\/.*Appointment`,
		onlyAdmin: false,
	},
	{
		regex:     `\/fitness_aggregator\.v1\.FitnessAggregator\/Get(Trainers?|Studios?|Class(es)?|Client)`,
		onlyAdmin: false,
	},
	{
		regex:     `\/fitness_aggregator\.v1\.ExampleService\/*`,
		onlyAdmin: false,
	},
	{
		regex:     `\/fitness_aggregator\.v1\.FitnessAggregator\/*`,
		onlyAdmin: true,
	},
	{
		regex:     `\/fitness_aggregator\.v1\.AdminPanel\/*`,
		onlyAdmin: true,
	},
}

func (a *Authorizer) PermissionInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	userID := GetUserIDFromContext(ctx)

	// Check if endpoint is `onlyAdmin` and user is not admin
	for _, validator := range endpointPermissionValidators {
		if !regexp.MustCompile(validator.regex).MatchString(info.FullMethod) {
			continue
		}

		slog.Info(fmt.Sprintf("Matched with regex: %s", validator.regex))

		if !validator.onlyAdmin {
			break
		}

		if userID != adminID {
			return nil, status.Error(codes.PermissionDenied, "admin permission required")
		}
	}

	return handler(ctx, req)
}

func GetUserIDFromContext(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}

	userIDs, ok := md[userIDMDKey]
	if !ok || len(userIDs) == 0 {
		return ""
	}

	return userIDs[0]
}

func (a *Authorizer) Auth(
	ctx context.Context, req *gen.AuthRequest,
) (*gen.AuthResponse, error) {
	id, err := a.Repo.GetIDByCreds(ctx, req.Phone, req.Password)
	switch {
	case errors.Is(err, db.ErrNotFound):
		return nil, status.Error(codes.NotFound, "user not found")
	case errors.Is(err, db.ErrInvalidCredentials):
		return nil, status.Error(codes.Unauthenticated, "invalid credentials")
	case err != nil:
		return nil, status.Error(codes.Internal, err.Error())
	}

	claims := &Claims{
		UserID: id.Hex(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(a.JWTSecret))
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &gen.AuthResponse{Token: tokenString, Id: id.Hex()}, nil
}

func (a *Authorizer) IsAdmin(_ context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (a *Authorizer) AuthInterceptor(
	ctx context.Context,
	req interface{},
	_ *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		slog.Info("metadata not found")
		return handler(ctx, req)
	}

	cookiesMetadata, ok := md[gRPCGatewayCookie]
	if !ok || len(cookiesMetadata) == 0 {
		slog.Info("authorization cookie not found")
		return handler(ctx, req)
	}

	rawCookies := cookiesMetadata[0]
	cookies, err := http.ParseCookie(rawCookies)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid cookies: %v", err)
	}

	cookieMap := map[string]string{}
	for _, c := range cookies {
		cookieMap[c.Name] = c.Value
	}

	token, ok := cookieMap[authorizationCookieKey]
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "authorization cookie not found")
	}

	if token == a.AdminToken {
		md.Append(userIDMDKey, adminID)
	} else {
		claims, err := validateJWT(token, a.JWTSecret)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, err.Error())
		}
		md.Set(userIDMDKey, claims.UserID)
	}

	ctx = metadata.NewIncomingContext(ctx, md)
	return handler(ctx, req)
}

func validateJWT(inputToken, jwtSecret string) (Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(inputToken, claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		})
	if err != nil {
		return Claims{}, err
	}

	if !token.Valid {
		return Claims{}, fmt.Errorf("invalid token")
	}

	if time.Now().Unix() > claims.ExpiresAt {
		return Claims{}, fmt.Errorf("token expired")
	}

	return *claims, nil
}

func checkUserToTargetPermissions(ctx context.Context, targetID string) error {
	switch GetUserIDFromContext(ctx) {
	case adminID, targetID:
		return nil
	default:
		return permissionDeniedErr
	}
}