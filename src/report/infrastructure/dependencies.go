package infraestructure

import (
	"apidos/src/core"
	"apidos/src/report/application/usescases"
	"apidos/src/report/infrastructure/adapters"
	"log"
)

type Dependencies struct {
	ProcessReportUseCase *usescases.ProcessReport
}

func NewDependencies(db *MySQL) (*Dependencies, error) {
	err := core.InitRabbitMQ()
	if err != nil {
		log.Fatal("Error al inicializar RabbitMQ", err)
	}

	rabbitService := adapters.NewRabbitMQService()
	rabbitPublishService := adapters.NewRabbitMQPublishService()
	reportRepo := NewMySQL()

	processReportUseCase := usescases.NewMysqlProcessReport(reportRepo, rabbitService, rabbitPublishService)

	go processReportUseCase.StartProcessingReports()
	return &Dependencies{
		ProcessReportUseCase: processReportUseCase,
	}, nil
}
