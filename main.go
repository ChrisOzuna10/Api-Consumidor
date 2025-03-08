package main

import (
	infraestructure "apidos/src/report/infrastructure"
	"apidos/src/report/infrastructure/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db := infraestructure.NewMySQL()

	deps, err := infraestructure.NewDependencies(db)
	if err != nil {
		log.Fatal("Error al inicializar dependencias:", err)
	}

	r := gin.Default()

	routes.RegisterRoutes(r, deps.ProcessReportUseCase)

	r.Run(":8081")
}
