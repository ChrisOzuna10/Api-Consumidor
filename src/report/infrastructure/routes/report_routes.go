package routes

import (
	"apidos/src/report/infrastructure/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigureRoutesReport(
	r *gin.Engine,
	processReportController *controllers.ProcessReportController,
) {
	// Ruta para procesar los reportes
	r.POST("/process-report", processReportController.Execute)
}
