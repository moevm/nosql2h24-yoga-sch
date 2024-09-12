package server

import (
	"log"
	"net"

	genv1 "gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/gen/proto/v1"
	v1 "gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/handlers/v1"

	"google.golang.org/grpc"
)

func New() {
	grpcServer := grpc.NewServer()

	genv1.RegisterLibraryServiceServer(grpcServer, &v1.LibraryServiceServer{})

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("error listening port :8080", err)
	}

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal("error serving port :8080", err)
	}
}
