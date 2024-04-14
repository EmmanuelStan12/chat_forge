package main

import (
	"context"
	"fmt"
	"os"

	"github.com/EmmanuelStan12/chat_forge/internal/app"
)

func main() {
	ctx := context.Background()
	if err := app.Run(ctx, os.Stdout, os.Args); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}
