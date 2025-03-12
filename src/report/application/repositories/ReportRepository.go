package repositories

import (
	"apidos/src/report/domain/entities"
	"log"
)

type ServiceNotification struct {
	musicService IMessageService
}

func NewServiceNotification(musicService IMessageService) *ServiceNotification {
	return &ServiceNotification{musicService: musicService}
}

func (sn *ServiceNotification) PublishEvent(eventType string, music entities.Notification) error {
	log.Println("Publishing event:", eventType)

	err := sn.musicService.PublishEvent(eventType, music)
	if err != nil {
		log.Printf("Error publishing event: %v", err)
		return err
	}
	return nil
}
