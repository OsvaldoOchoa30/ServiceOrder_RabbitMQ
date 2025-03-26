package repositories

import (
    "service_order/src/reservation/domain/entities"
    "time"
)

type IReservation interface {
    Save(username string, lastname string, cellphone string, email string, reservationdate time.Time, status string, numbercard string, pin int64) (*entities.Reservation, error)
}