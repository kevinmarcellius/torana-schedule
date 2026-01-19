package repository

import (
	"github.com/kevinmarcellius/torana-schedule/internal/model"
	"gorm.io/gorm"
)

type LineRepository struct {
	db *gorm.DB
}

func NewLineRepository(db *gorm.DB) *LineRepository {
	return &LineRepository{db: db}
}

func (r *LineRepository) GetLinesAndStations() ([]model.TrainTrip, error) {
	var trips []model.TrainTrip
	result := r.db.Table("train_trips").
		Select("line, station, distance, train_type").
		Order("line, distance ASC").
		Find(&trips)

	if result.Error != nil {
		return nil, result.Error
	}
	return trips, nil
}
