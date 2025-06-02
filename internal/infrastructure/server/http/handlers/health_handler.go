package handlers

import (
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HealthHandler struct {
	db *gorm.DB
}

func NewHealthHandler(db *gorm.DB) *HealthHandler {
	return &HealthHandler{
		db: db,
	}
}

func (h *HealthHandler) CheckHealth(c *gin.Context) {
	var dbStatus = "up"
	sqlDB, err := h.db.DB()
	if err != nil {
		dbStatus = "down"
	} else if err := sqlDB.Ping(); err != nil {
		dbStatus = "down"
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"info": gin.H{
			"database_status": dbStatus,
			"go_version":      runtime.Version(),
			"go_os":           runtime.GOOS,
			"go_arch":         runtime.GOARCH,
			"cpu_cores":       runtime.NumCPU(),
			"goroutines":      runtime.NumGoroutine(),
		},
	})
}
