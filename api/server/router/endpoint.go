package router

import (
	"bottleneck-lab/server/handler"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func RegisterEndpoints(r *gin.Engine, db *sql.DB) {
	healthHandler := handler.NewHealthHandler(db)
	metricsHandler := handler.NewMetricsHandler(db)
	queryHandler := handler.NewQueryHandler(db)

	r.GET("/health", healthHandler.Check)
	r.GET("/metrics", metricsHandler.Get)
	r.GET("/slow-query", queryHandler.SlowQuery)
	r.GET("/heavy", queryHandler.Heavy)
}
