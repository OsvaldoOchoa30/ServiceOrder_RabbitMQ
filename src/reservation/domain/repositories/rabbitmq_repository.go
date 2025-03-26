package repositories

import (
	"service_order/src/reservation/domain/entities"
	_"time"
)

type IReservationRabbitqm interface {
    Save(reservation *entities.Reservation) error
}

//regla de dependecia
//patron repositorio