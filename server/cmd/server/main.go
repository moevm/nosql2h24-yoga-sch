package main

import (
	"flag"

	"gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/server"
)

var (
	httpPort = flag.Int("http-port", 4200, "The server port")
	grpcPort = flag.Int("grpc-port", 4201, "The server port")
)

func main() {
	flag.Parse()

	server.Run(*httpPort, *grpcPort)
}
