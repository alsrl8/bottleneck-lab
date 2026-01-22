package router

import (
	"bottleneck-lab/server/handler"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterEndpoints(r *gin.Engine, db *sql.DB) {
	metricsHandler := handler.NewMetricsHandler(db)

	r.GET("/metrics", metricsHandler.Get)
}
