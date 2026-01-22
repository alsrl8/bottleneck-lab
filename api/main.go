package main

import (
	"bottleneck-lab/db"
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	slog.Info("bottleneck-lab api server started")

	d, err := db.Connect()
	if err != nil {
		slog.Error("failed to connect to database", "err", err)
		panic(err)
	}
	slog.Info("database connected")

	// Health check endpoint
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		// Check database connection
		if err := d.Ping(); err != nil {
			slog.Error("health check failed: database ping error", "err", err)
			w.WriteHeader(http.StatusServiceUnavailable)
			_, _ = w.Write([]byte("Service Unavailable: DB connection error"))
			return
		}

		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("OK"))
		if err != nil {
			slog.Error("failed to write health check response", "err", err)
			return
		}
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: http.DefaultServeMux,
	}

	go func() {
		slog.Info("starting http server", "addr", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("http server error", "err", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("shutting down server...")

	// Graceful shutdown with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		slog.Error("server forced to shutdown", "err", err)
	}

	if err := d.Close(); err != nil {
		slog.Error("failed to close database connection", "err", err)
	} else {
		slog.Info("database connection closed")
	}

	slog.Info("bottleneck-lab api server stopped")
}
