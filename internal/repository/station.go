package repository

import (
	"github.com/kevinmarcellius/torana-schedule/internal/model"
	"gorm.io/gorm"
)

type StationRepository struct {
	db *gorm.DB
}

func NewStationRepository(db *gorm.DB) *StationRepository {
	return &StationRepository{db: db}
}

func (r *StationRepository) GetStationsAndSchedulesByLine(lineName string) ([]model.StationScheduleData, error) {
	var data []model.StationScheduleData

	result := r.db.Table("train_trips tt").
		Select("tt.station as station_name, tt.distance, to_char(s.time, 'HH24:MI:SS') as scheduled_time").
		Joins("JOIN schedules s ON s.trip_id = tt.id").
		Where("tt.line = ?", lineName).
		Order("tt.distance ASC, scheduled_time ASC").
		Find(&data)

	if result.Error != nil {
		return nil, result.Error
	}

	return data, nil
}
