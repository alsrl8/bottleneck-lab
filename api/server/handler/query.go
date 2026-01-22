package handler

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type QueryHandler struct {
	DB *sql.DB
}

func NewQueryHandler(db *sql.DB) *QueryHandler {
	return &QueryHandler{DB: db}
}

func (h *QueryHandler) SlowQuery(c *gin.Context) {
	var result int
	err := h.DB.QueryRow("SELECT SLEEP(0.1)").Scan(&result)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "ok"})
}
