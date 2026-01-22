package handler

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthHandler struct {
	db *sql.DB
}

func NewHealthHandler(db *sql.DB) *HealthHandler {
	return &HealthHandler{db: db}
}

func (h *HealthHandler) Check(c *gin.Context) {
	if err := h.db.Ping(); err != nil {
		slog.Error("health check failed: database ping error", "err", err)
		c.String(http.StatusServiceUnavailable, "Service Unavailable: DB connection error")
		return
	}

	c.String(http.StatusOK, "OK")
}
