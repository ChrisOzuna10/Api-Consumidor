package controllers

import (
	"apidos/src/report/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ReportController struct {
	repo domain.IReport
}

func NewReportController(repo domain.IReport) *ReportController {
	return &ReportController{repo: repo}
}

func (ctrl *ReportController) GetAllReports(c *gin.Context) {
	alerts, err := ctrl.repo.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error obteniendo los reportes"})
		return
	}
	c.JSON(http.StatusOK, alerts)
}
