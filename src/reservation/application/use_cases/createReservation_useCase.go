package usecases

import (
    "log"
    "service_order/src/reservation/domain/entities"
    "service_order/src/reservation/domain/repositories"
    "time"
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

func (cp *CreateReservationUseCase) Execute(username string, lastname string, cellphone string, email string, reservationdate time.Time, status string, numbercard string, pin int64) (*entities.Reservation, error) {
    reservation, err := cp.db.Save(username, lastname, cellphone, email, reservationdate, status, numbercard, pin)
    if err != nil {
        log.Printf("Error al guardar la reservación: %s", err)
        return nil, err
    }
    
    errSendMessage := cp.rtqm.Save(reservation)
    if errSendMessage != nil {
        log.Printf("Error al enviar mensaje: %s", errSendMessage)
        return reservation, errSendMessage
    }
    
    return reservation, nil
}