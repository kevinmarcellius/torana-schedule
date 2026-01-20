package repository

import (
	"github.com/kevinmarcellius/torana-schedule/internal/model"
	"gorm.io/gorm"
)

type TripRepository interface {
	GetAllTrips() ([]model.TrainTrip, error)
}

type tripRepository struct {
	db *gorm.DB
}

func NewTripRepository(db *gorm.DB) TripRepository {
	return &tripRepository{db: db}
}

func (r *tripRepository) GetAllTrips() ([]model.TrainTrip, error) {
	var trips []model.TrainTrip
	if err := r.db.Table("train_trips").Find(&trips).Error; err != nil {
		return nil, err
	}
	return trips, nil
}
