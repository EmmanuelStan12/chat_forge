package app

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
)

func Run(ctx context.Context, w io.Writer, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()
	fmt.Println("Hello from our first web application in golang")
	return nil
}
