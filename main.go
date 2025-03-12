package main

import (
	"apidos/src/report/infrastructure"
	"apidos/src/report/infrastructure/routes"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Inicializamos las dependencias para el procesamiento de reportes
	dependencies := infraestructure.InitReport()
	defer dependencies.RabbitMQAdapter.Close()

	// Creamos una nueva instancia de Gin
	r := gin.Default()

	// Configuramos los encabezados CORS
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// Registramos las rutas para el reporte
	routes.ConfigureRoutesReport(r, dependencies.CreateMusicController)

	// Iniciamos el servidor en el puerto 8081
	if err := r.Run(":8088"); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
