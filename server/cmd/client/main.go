package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	v1 "gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/gen/proto/v1"
)

func main() {
	grpcOpt := grpc.WithTransportCredentials(insecure.NewCredentials())
	grpcC, err := grpc.NewClient(":80", grpcOpt)
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to create grpc client: %v", err))
	}

	defer func() {
		if err := grpcC.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	c := v1.NewExampleServiceClient(grpcC)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	resp, err := c.Echo(ctx, &v1.EchoRequest{Value: "zeleboba"})
	if err != nil {
		log.Fatal(fmt.Sprintf("Failed to echo: %v", err))
	}

	log.Println(resp.GetValue())
}
