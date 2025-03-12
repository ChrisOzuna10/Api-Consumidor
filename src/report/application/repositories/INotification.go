package repositories

import "apidos/src/report/domain/entities"

type IMessageService interface {
	PublishEvent(eventType string, musics entities.Notification) error
}
