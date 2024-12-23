package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/db"
	genv1 "gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/gen/proto/v1"
	v1 "gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/handlers/v1"
	"gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/middlewares"
)

func newGRPC(
	port int,
	jwtSecret, adminToken string,
	mgClient *mongo.Client,
) (net.Listener, *grpc.Server) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("error listening address %d: %v", port, err)
	}

	repo := db.NewMongoRepository(mgClient)

	authorizer := &v1.Authorizer{
		Repo:       repo,
		JWTSecret:  jwtSecret,
		AdminToken: adminToken,
	}

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			authorizer.AuthInterceptor,
			authorizer.PermissionInterceptor,
		),
	)

	genv1.RegisterAuthorizerServer(s, authorizer)
	genv1.RegisterAdminPanelServer(s, &v1.AdminPanel{Repo: repo})
	genv1.RegisterSearchEngineServer(s, &v1.SearchEngine{Repo: repo})
	genv1.RegisterFitnessAggregatorServer(s, &v1.FitnessAggregator{Repo: repo})
	genv1.RegisterExampleServiceServer(s, &v1.ExampleService{DbC: mgClient})

	fmt.Printf("Starting grpc server on port %d...\n", port)
	return l, s
}

type HandlerRegistrar = func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error

func newHTTP(httpPort, grpcPort int) *http.Server {
	mux := runtime.NewServeMux()

	grpcOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	handlersToRegister := []HandlerRegistrar{
		genv1.RegisterAuthorizerHandlerFromEndpoint,
		genv1.RegisterAdminPanelHandlerFromEndpoint,
		genv1.RegisterSearchEngineHandlerFromEndpoint,
		genv1.RegisterFitnessAggregatorHandlerFromEndpoint,
		genv1.RegisterExampleServiceHandlerFromEndpoint,
	}

	for _, h := range handlersToRegister {
		if err := h(
			context.Background(), mux, fmt.Sprintf(":%d", grpcPort), grpcOpts,
		); err != nil {
			log.Fatalf("failed to register http endpoint: %v", err)
		}
	}

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", httpPort),
		Handler: middlewares.WithLogger(mux),
	}

	fmt.Printf("Starting http server on port %d...\n", httpPort)
	return s
}