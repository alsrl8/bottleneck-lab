package router

import (
	"bottleneck-lab/server/handler"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterHealth(r *gin.Engine, db *sql.DB) {
	h := handler.NewHealthHandler(db)
	r.GET("/health", h.Check)
}
