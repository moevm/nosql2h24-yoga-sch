package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/credentials/insecure"

	genv1 "gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/gen/proto/v1"
	v1 "gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/handlers/v1"
	"gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/middlewares"

	"google.golang.org/grpc"
)

func Run(httpPort, grpcPort int) {
	l, grpcS := newGRPC(grpcPort)
	httpS := newHTTP(httpPort, grpcPort)

	fmt.Printf("Starting http server on port %d...\n", httpPort)
	fmt.Printf("Starting grpc server on port %d...\n", grpcPort)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error { return grpcS.Serve(l) })
	g.Go(func() error { return httpS.ListenAndServe() })
	g.Go(func() error {
		<-gCtx.Done()
		fmt.Println("Shutting down server")

		err := httpS.Shutdown(ctx)
		grpcS.GracefulStop()

		return err
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("Exit reason: %v\n", err)
	}
}

func newGRPC(port int) (net.Listener, *grpc.Server) {
	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("error listening address %d: %v", port, err)
	}

	s := grpc.NewServer()
	genv1.RegisterPlaceRepositoryServer(s, &v1.PlaceRepository{})

	return l, s
}

func newHTTP(httpPort, grpcPort int) *http.Server {
	mux := runtime.NewServeMux()

	grpcOpts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	if err := genv1.RegisterPlaceRepositoryHandlerFromEndpoint(
		context.Background(), mux, fmt.Sprintf(":%d", grpcPort), grpcOpts,
	); err != nil {
		log.Fatalf("failed to register http endpoint: %v", err)
	}

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", httpPort),
		Handler: middlewares.EnableCors(middlewares.WithLogger(mux)),
	}

	return s
}
