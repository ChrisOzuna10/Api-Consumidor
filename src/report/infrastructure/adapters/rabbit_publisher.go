package adapters

import (
	"apidos/src/core"
	"apidos/src/report/domain/entities"
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitPublishService struct{}

func NewRabbitMQPublishService() *RabbitPublishService {
	return &RabbitPublishService{}
}

func (s *RabbitPublishService) PublishReport(report *entities.Report) error {
	if core.RabbitChannel == nil {
		log.Println("No se conecto a RabbitMQ channel")
		return nil
	}

	body, err := json.Marshal(report)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = core.RabbitChannel.PublishWithContext(
		ctx,
		"report",
		"process",
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)

	if err != nil {
		log.Println("Error al enviar reporte", err)
	} else {
		log.Println("Se envio el reporte")
	}
	return err
}
