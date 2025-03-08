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

func (repo *MySQLRepository) Save(username string, lastname string, cellphone string, email string, reservationdate time.Time, status string) error {
	reservation := entities.NewReservation(username, lastname, cellphone, email, reservationdate, status)
	result := repo.db.Create(&reservation)
	if result.Error != nil {
		return result.Error
	}
	return nil
}