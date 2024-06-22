package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bohdanch-w/parallel/cmd/parallel"
)

func main() {
	ctx := OSInterruptContext(context.Background())

	if err := parallel.RunParallel(ctx, os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "execute parallel: %v\n", err)
	}
}

func OSInterruptContext(c context.Context) context.Context {
	ctx, cancel := context.WithCancel(c)

	go func() {
		defer cancel()

		shutdown := make(chan os.Signal, 1)
		signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

		defer signal.Stop(shutdown)

		select {
		case <-shutdown:
		case <-ctx.Done():
		}
	}()

	return ctx
}
