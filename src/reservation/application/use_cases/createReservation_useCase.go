package usecases

import (
	"service_order/src/reservation/domain/repositories"
	"time"
	"log"
)

type CreateReservationUseCase struct {
	db repositories.IReservation
	rtqm repositories.IReservationRabbitqm
}

func NewCreateReservationUseCase(db repositories.IReservation, rtqm repositories.IReservationRabbitqm) *CreateReservationUseCase {
	return &CreateReservationUseCase{db: db, rtqm: rtqm}
}

func (usecase *CreateReservationUseCase) SetOrder(mysqlRepository repositories.IReservation, rabbitqmRepository repositories.IReservationRabbitqm) {
	usecase.db = mysqlRepository
	usecase.rtqm = rabbitqmRepository
}


func (cp *CreateReservationUseCase) Execute(username string, lastname string, cellphone string, email string, reservationdate time.Time, status string) error {
	err := cp.db.Save(username, lastname, cellphone, email, reservationdate, status)
	errSendMessage := cp.rtqm.Save(username, lastname, cellphone, email, reservationdate, status)
	if err != nil || errSendMessage != nil {
		log.Panicf("error to send message %s", err); 
		return err
	}
	
	return errSendMessage; 
}
