package usescases

import (
	"apidos/src/report/domain"
	"apidos/src/report/domain/entities"
	"apidos/src/report/infrastructure/adapters"
	"log"
	"time"
)

type ProcessReport struct {
	Repo           domain.IReport
	Rabbit         *adapters.RabbitMQService
	PublishService *adapters.RabbitPublishService
}

func NewMysqlProcessReport(repo domain.IReport, rabbit *adapters.RabbitMQService, publishService *adapters.RabbitPublishService) *ProcessReport {
	return &ProcessReport{
		Repo:           repo,
		Rabbit:         rabbit,
		PublishService: publishService,
	}
}

func (pr *ProcessReport) StartProcessingReports() {
	for {
		time.Sleep(5 * time.Second)

		reports, err := pr.Rabbit.FetchReport()
		if err != nil {
			log.Println("Error obteniendo reportes pendientes:", err)
			continue
		}

		for _, data := range reports {
			// Convertir el mapa a una estructura Report
			report := entities.Report{
				ID:      int(data["id"].(float64)), // JSON lo devuelve como float64
				Title:   data["title"].(string),
				Content: data["content"].(string),
				Status:  "processed",
			}

			log.Printf("Procesando reporte ID %d", report.ID)

			// Actualizar el reporte en la base de datos
			err := pr.Repo.Update(report.ID, report.Title, report.Content, report.Status)
			if err != nil {
				log.Println("Error actualizando el reporte:", err)
				continue
			}

			// Enviar mensaje a la cola de RabbitMQ
			err = pr.PublishService.PublishReport(&report)
			if err != nil {
				log.Println("Error enviando mensaje a RabbitMQ:", err)
			} else {
				log.Printf("Reporte ID %d procesado y enviado a la cola", report.ID)
			}
		}
	}
}
