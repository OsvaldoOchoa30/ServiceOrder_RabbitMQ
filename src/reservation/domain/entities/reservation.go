package entities

import "time"

type Reservation struct {
	Id              int32 `json:"id" gorm:"column:id_reservation;primaryKey;autoIncrement"`
	UserName        string `json:"username"`
	LastName        string `json:"lastname"`
	CellPhone       string `json:"cellphone"`
	Email           string `json:"email"`
	ReservationDate time.Time `json:"reservationdate"`
	Status          string `json:"status"`
}

func NewReservation(username string, lastname string, cellphone string, email string, reservationdate time.Time, status string) *Reservation {
	return &Reservation{
		UserName: username,
		LastName: lastname,
		CellPhone: cellphone,
		Email: email,
		ReservationDate: reservationdate,
		Status: status,
	}
}