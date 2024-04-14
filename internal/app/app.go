package app

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Run(ctx context.Context, w io.Writer, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	logger := &log.Logger{}
	router := NewServerHandler(logger)
	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	go func() {
		log.Printf("listening on %s\n", server.Addr)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(os.Stderr, "Error listening and serving: %s\n", err)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done() // Wait for the context to be canceled

		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(os.Stderr, "error shutting down server: %s\n", err)
		}
	}()

	wg.Wait()
	return nil
}

// The NewServerHandler constructor is responsible for all the top-level HTTP stuff that applies to all endpoints, like CORS, auth middleware, and logging
func NewServerHandler(logger *log.Logger) http.Handler {
	handler := chi.NewRouter()

	// Initialize base middlewares
	handler.Use(middleware.Logger)
	handler.Use(middleware.RequestID)

	// Initialize routes

	return handler
}

func indexHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello world"))
}
