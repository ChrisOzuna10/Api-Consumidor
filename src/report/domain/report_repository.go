package domain

import "apidos/src/report/domain/entities"

// Interfaz para el repositorio que maneja las notificaciones
type INotification interface {
	Send(music entities.Notification) error
}
