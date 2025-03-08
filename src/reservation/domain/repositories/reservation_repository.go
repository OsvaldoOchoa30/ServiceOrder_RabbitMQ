package repositories

import (
	"time"
)

type IReservation interface {
	Save(username string, lastname string, cellphone string, email string, reservationdate time.Time,status string) error
}

