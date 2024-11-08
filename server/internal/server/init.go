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

	genv1 "gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/gen/proto/v1"
	v1 "gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/handlers/v1"
	"gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/middlewares"
)

func newGRPC(port int, mongoClient *mongo.Client) (net.Listener, *grpc.Server) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("error listening address %d: %v", port, err)
	}

	s := grpc.NewServer()
	genv1.RegisterExampleServiceServer(s, &v1.ExampleService{DbC: mongoClient})

	fmt.Printf("Starting grpc server on port %d...\n", port)
	return l, s
}

func newHTTP(httpPort, grpcPort int) *http.Server {
	mux := runtime.NewServeMux()

	grpcOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := genv1.RegisterExampleServiceHandlerFromEndpoint(
		context.Background(), mux, fmt.Sprintf(":%d", grpcPort), grpcOpts,
	); err != nil {
		log.Fatalf("failed to register http endpoint: %v", err)
	}

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", httpPort),
		Handler: middlewares.WithLogger(mux),
	}

	fmt.Printf("Starting http server on port %d...\n", httpPort)
	return s
}
