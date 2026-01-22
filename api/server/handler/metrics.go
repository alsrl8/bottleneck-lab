package handler

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type MetricsHandler struct {
	db *sql.DB
}

func NewMetricsHandler(db *sql.DB) *MetricsHandler {
	return &MetricsHandler{db: db}
}

func (h *MetricsHandler) Get(c *gin.Context) {
	stats := h.db.Stats()
	c.JSON(200, stats)
}
