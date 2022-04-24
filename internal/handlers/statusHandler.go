package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zepyrshut/go-web-starter-gin/internal/config"
)

func (m *Repository) GetStatusHandler(c *gin.Context) {

	currentStatus := config.AppStatus{
		Status:      "OK",
		Environment: m.App.Status.Environment,
		Version:     m.App.Status.Version,
	}

	c.Header("Content-Type", "application/json")
	c.Status(http.StatusOK)
	c.JSON(http.StatusOK, currentStatus)
}
