package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/AnnDutova/in-memory-db/internal/service"
)

func main() {
	fmt.Println(os.Getpid())
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	svc, err := service.NewService()
	if err != nil {
		panic(err)
	}

	errChan := make(chan error, 1)

	go func() {
		err = svc.Run(ctx)
		if err != nil {
			errChan <- err
		}
	}()

	select {
	case <-ctx.Done():
		fmt.Fprint(os.Stdin, "execution ended")
	case err := <-errChan:
		fmt.Fprint(os.Stderr, err.Error())
	}
}
