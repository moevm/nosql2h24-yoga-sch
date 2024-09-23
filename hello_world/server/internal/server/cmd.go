package server

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"

	"gitlab.com/purposeless-lab/monorepo/fitness-aggregator/internal/db"
)

func Run(httpPort, grpcPort int) {
	dbC := db.New()

	l, grpcS := newGRPC(grpcPort, dbC)
	httpS := newHTTP(httpPort, grpcPort)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	g, gCtx := errgroup.WithContext(ctx)
	g.Go(func() error { return grpcS.Serve(l) })
	g.Go(func() error { return httpS.ListenAndServe() })
	g.Go(func() error {
		<-gCtx.Done()
		fmt.Println("Shutting down server")

		httpErr := httpS.Shutdown(context.Background())
		grpcS.GracefulStop()

		mongoErr := dbC.Disconnect(context.Background())

		return errors.Join(httpErr, mongoErr)
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("Exit reason: %v\n", err)
	}
	fmt.Println("Exiting")
}
