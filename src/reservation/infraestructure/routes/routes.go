package routes

import (
	"service_order/src/reservation/application/use_cases"
	"service_order/src/reservation/infraestructure/controllers"
	"service_order/src/reservation/domain/repositories"
	"github.com/gin-gonic/gin"
)

func SetupReservationsRoutes(
	router *gin.Engine, 
	repo repositories.IReservation, 
	rabbitRepo repositories.IReservationRabbitqm, // Nombre de interfaz corregido
	
) {
	createReservationsCaseUse := usecases.NewCreateReservationUseCase(repo, rabbitRepo)
	createReservationsController := controllers.NewCreateReservationController(createReservationsCaseUse) // Nombre corregido

	reservationGroup := router.Group("/reservations")
	{
		reservationGroup.POST("", createReservationsController.Run)
	}
}