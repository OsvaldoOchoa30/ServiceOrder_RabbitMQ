package repositories

import "time"

type IReservationRabbitqm interface {
    Save(username, lastname, cellphone, email string, reservationdate time.Time, status string) error
}