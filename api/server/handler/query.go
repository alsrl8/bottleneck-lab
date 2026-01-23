package handler

import (
	d "bottleneck-lab/db"
	"database/sql"
	"log/slog"

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

func (h *QueryHandler) Heavy(c *gin.Context) {
	results, err := h.loadAllData(h.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"status": "ok", "data": results})
}

func (h *QueryHandler) loadAllData(db *sql.DB) ([]d.SensorLog, error) {
	rows, err := db.Query("SELECT * FROM sensor_logs")
	if err != nil {
		slog.Error("failed to query sensor logs", "err", err)
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	var results []d.SensorLog
	for rows.Next() {
		var log d.SensorLog
		err = rows.Scan(&log.Id, &log.SensorId, &log.Value, &log.CreatedAt)
		if err != nil {
			slog.Error("failed to scan sensor log", "err", err)
			continue
		}
		results = append(results, log)
	}
	return results, nil
}
