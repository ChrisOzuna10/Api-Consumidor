package application

import (
	"apidos/src/report/application/repositories"
	"apidos/src/report/domain"
	"apidos/src/report/domain/entities"
	"log"
)

type ProcessReport struct {
	reportRepo          domain.INotification
	serviceNotification repositories.IMessageService
}

func NewProcessReport(reportRepo domain.INotification, serviceNotification repositories.IMessageService) *ProcessReport {
	return &ProcessReport{
		reportRepo:          reportRepo,
		serviceNotification: serviceNotification,
	}
}

func (pr *ProcessReport) Execute(music entities.Notification) error {
	// Enviar la notificación al repositorio
	err := pr.reportRepo.Send(music)
	if err != nil {
		log.Println("Error enviando la notificación:", err)
		return err
	}

	// Publicar el evento sobre la notificación creada
	err = pr.serviceNotification.PublishEvent("NotificationCreated", music)
	if err != nil {
		log.Printf("Error notificando sobre la notificación creada: %v", err)
		return err
	}

	log.Printf("Notificación ID %d procesada y notificada", music.ID)
	return nil
}
