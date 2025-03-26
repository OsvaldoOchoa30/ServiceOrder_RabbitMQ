package adapters

import (
	"service_order/src/reservation/domain/entities"
	"service_order/src/reservation/domain/repositories"
	"time"
	"gorm.io/gorm"
)

type MySQLRepository struct {
	db *gorm.DB
}

func NewMySQLRepository(db *gorm.DB) repositories.IReservation {
	return &MySQLRepository{db: db}
}

func (repo *MySQLRepository) Save(username string, lastname string, cellphone string, email string, reservationdate time.Time, status string, numbercard string, pin int64) (*entities.Reservation, error) {
    reservation := entities.NewReservation(username, lastname, cellphone, email, reservationdate, status, numbercard, pin)
    result := repo.db.Create(&reservation)
    if result.Error != nil {
        return nil, result.Error
    }
    // Ahora reservation.Id debería tener el valor asignado por GORM
    return reservation, nil
}