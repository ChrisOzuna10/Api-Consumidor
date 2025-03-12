package infraestructure

import (
	"apidos/src/report/application/repositories"
	application "apidos/src/report/application/usescases"
	"apidos/src/report/infrastructure/adapters"
	"apidos/src/report/infrastructure/controllers"
	"log"
)

type DependenciesReport struct {
	CreateMusicController *controllers.ProcessReportController
	RabbitMQAdapter       *adapters.RabbitMQAdapter
}

func InitReport() *DependenciesReport {
	// Inicializamos RabbitMQ
	rmqClient, err := adapters.NewRabbitMQAdapter()
	// Pasa la URL de conexión si es necesario
	if err != nil {
		log.Fatalf("Error creando RabbitMQ cliente: %v", err)
	}

	// Creamos el servicio de notificación
	reportService := repositories.NewServiceNotification(rmqClient)

	// Creamos el caso de uso de procesamiento de reportes
	createMusic := application.NewProcessReport(rmqClient, reportService)

	return &DependenciesReport{
		CreateMusicController: controllers.NewProcessReportController(createMusic, rmqClient),
		RabbitMQAdapter:       rmqClient,
	}
}
