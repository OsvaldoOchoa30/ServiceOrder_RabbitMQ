package controllers

import (
	"service_order/src/reservation/application/use_cases"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

type CreateReservationController struct {
	useCase *usecases.CreateReservationUseCase
}

func NewCreateReservationController(useCase *usecases.CreateReservationUseCase) *CreateReservationController{
	return &CreateReservationController{useCase: useCase}
}

func (cp * CreateReservationController) Run(c * gin.Context) {
	var input struct {
		Id              int32
		UserName        string
		LastName        string
		CellPhone       string
		Email           string
		ReservationDate time.Time
		Status          string
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	err := cp.useCase.Execute(input.UserName, input.LastName, input.CellPhone, input.Email, input.ReservationDate, input.Status)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(http.StatusCreated, gin.H{"message": "Reservation created successfully"})
 

}