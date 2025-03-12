package controllers

import (
	"apidos/src/report/application/usescases"
	"apidos/src/report/domain"
	"apidos/src/report/domain/entities"
	"github.com/gin-gonic/gin"
)

type ProcessReportController struct {
	useCase *application.ProcessReport
	music   domain.INotification
}

func NewProcessReportController(useCase *application.ProcessReport, music domain.INotification) *ProcessReportController {
	return &ProcessReportController{useCase: useCase, music: music}
}

func (prc *ProcessReportController) Execute(c *gin.Context) {
	var music entities.Notification
	if err := c.ShouldBindJSON(&music); err != nil {
		c.JSON(400, gin.H{"error": "Datos inválidos"})
		return
	}

	err := prc.useCase.Execute(music)
	if err != nil {
		c.JSON(500, gin.H{"error": "No se pudo procesar el reporte"})
		return
	}
	c.JSON(200, gin.H{"message": "Musica creada correctamente", "Musica": music})
}
