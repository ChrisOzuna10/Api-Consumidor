package application

import (
	"apidos/src/report/domain/entities"
	"apidos/src/report/infrastructure/adapters"
)

type ProcessReport struct {
	rmqClient *adapters.RabbitMQAdapter
}

func NewProcessReport(rmqClient *adapters.RabbitMQAdapter) *ProcessReport {
	return &ProcessReport{rmqClient: rmqClient}
}

func (pr *ProcessReport) Send(music entities.Notification) error {
	return pr.rmqClient.Send(music)
}
