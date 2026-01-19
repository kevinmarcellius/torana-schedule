package repository

import (
	"github.com/kevinmarcellius/torana-schedule/internal/model"
	"gorm.io/gorm"
)

type TripRepository struct {
	db *gorm.DB
}

func NewTripRepository(db *gorm.DB) *TripRepository {
	return &TripRepository{db: db}
}

func (r *TripRepository) GetAllTrips() ([]model.TrainTrip, error) {
	var trips []model.TrainTrip
	if err := r.db.Table("train_trips").Find(&trips).Error; err != nil {
		return nil, err
	}
	return trips, nil
}
