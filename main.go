package main

import (
	"log"
	"service_order/src/core"
	"service_order/src/reservation/infraestructure/adapters"
	"service_order/src/reservation/infraestructure/routes" // Importar rutas
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func main() {
	db, err := core.ConnectToDataBase()
	if err != nil {
		log.Fatal("Error connecting to the database", err)
	}

	rabbitmq, err := core.GetChannel()
	if err != nil {
		log.Fatalf("Error al obtener la conexión a RabbitMQ: %v", err)
	}

	mysqlRepository := adapters.NewMySQLRepository(db)
	
	rabbitqmRepository := adapters.NewRabbitRepository(rabbitmq.Ch)
	
	// Ensure rabbitqmRepository implements the correct interface


	router := gin.Default()
	router.Use(cors.Default())

	// Registrar rutas
	routes.SetupReservationsRoutes(router, mysqlRepository, rabbitqmRepository) // Llamada añadida

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}