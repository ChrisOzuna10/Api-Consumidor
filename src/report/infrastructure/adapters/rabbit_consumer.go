package adapters

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"os"
)

type RabbitMQService struct{}

func NewRabbitMQService() *RabbitMQService {
	return &RabbitMQService{}
}

func (r *RabbitMQService) FetchReport() ([]map[string]interface{}, error) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	alertsURL := os.Getenv("REPORTS_API_URL")

	resp, err := http.Get(alertsURL)
	if err != nil {
		log.Println("Error fetching alerts from RabbitMQ", err)
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading alerts from RabbitMQ", err)
		return nil, err
	}

	var alerts []map[string]interface{}
	if err := json.Unmarshal(body, &alerts); err != nil {
		return nil, err
	}
	return alerts, nil
}
