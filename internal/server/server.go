package server

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/egomes/schedule/internal/env"
)

const shutdownTimeout = 15 * time.Second

func StartServer(envs *env.Env) error {
	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", envs.Port),
		Handler: mux,
	}

	serverErr := make(chan error, 1)

	go func() {
		log.Println("Server is running on", server.Addr)
		serverErr <- server.ListenAndServe()
	}()

	signalCtx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	defer stop()

	select {
	case err := <-serverErr:
		if !errors.Is(err, http.ErrServerClosed) {
			return err
		}
		return nil
	case <-signalCtx.Done():
		stop()
	}

	shutdownCtx, cancel := context.WithTimeout(
		context.Background(),
		shutdownTimeout,
	)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		closeErr := server.Close()
		return errors.Join(fmt.Errorf("Gracefull shutdown: %w", err), closeErr)
	}

	err := <-serverErr

	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return fmt.Errorf("Server HTTP: %w", err)
	}

	return nil
}
