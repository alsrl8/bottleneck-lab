package main

import (
	"bottleneck-lab/db"
	"bottleneck-lab/server/router"
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	slog.Info("bottleneck-lab api server started")

	d, err := db.Connect()
	if err != nil {
		slog.Error("failed to connect to database", "err", err)
		panic(err)
	}
	slog.Info("database connected")

	// Setup Gin router
	r := gin.Default()
	router.RegisterHealth(r, d)

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
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
