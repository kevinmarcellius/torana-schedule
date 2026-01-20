package repository

import (
	"github.com/kevinmarcellius/torana-schedule/internal/model"
	"gorm.io/gorm"
)

type LineRepository interface {
	GetLinesAndStations() ([]model.TrainTrip, error)
}


type lineRepository struct {
	db *gorm.DB
}

func NewLineRepository(db *gorm.DB) LineRepository {
	return &lineRepository{db: db}
}

func (r *lineRepository) GetLinesAndStations() ([]model.TrainTrip, error) {
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
