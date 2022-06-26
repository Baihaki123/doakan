package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type dashboardHandler struct {
}

func NewDashboardHandler() *dashboardHandler {
	return &dashboardHandler{}
}

func (h *dashboardHandler) Dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "dashboard.html", nil)
}
