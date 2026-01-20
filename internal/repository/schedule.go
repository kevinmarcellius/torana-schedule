package repository

import (
	"github.com/kevinmarcellius/torana-schedule/internal/model"
	"gorm.io/gorm"
)

type ScheduleRepository interface{
	GetAllSchedules() ([]model.Schedule, error)
	SearchSchedules(params *model.ScheduleSearchParams) ([]model.Schedule, error)
}

type scheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &scheduleRepository{db: db}
}

// GetAllSchedules retrieves all schedules, ordered by station and time.
func (r *scheduleRepository) GetAllSchedules() ([]model.Schedule, error) {
	var schedules []model.Schedule
	result := r.db.Table("schedules s").
		Select("tt.line, tt.station, tt.train_type, to_char(s.time, 'HH24:MI:SS') as time").
		Joins("JOIN train_trips tt ON tt.id = s.trip_id").
		Order("tt.station, s.time ASC").
		Find(&schedules)

	if result.Error != nil {
		return nil, result.Error
	}
	return schedules, nil
}

// SearchSchedules retrieves schedules based on optional search criteria.
func (r *scheduleRepository) SearchSchedules(params *model.ScheduleSearchParams) ([]model.Schedule, error) {
	var schedules []model.Schedule

	query := r.db.Table("schedules s").
		Select("tt.line, tt.station, tt.train_type, to_char(s.time, 'HH24:MI:SS') as time").
		Joins("JOIN train_trips tt ON tt.id = s.trip_id")

	if params.Station != "" {
		query = query.Where("tt.station = ?", params.Station)
	}
	if !params.StartTime.IsZero() {
		query = query.Where("s.time >= ?", params.StartTime)
	}
	if !params.EndTime.IsZero() {
		query = query.Where("s.time <= ?", params.EndTime)
	}
	if params.Line != "" {
		query = query.Where("tt.line = ?", params.Line)
	}
	if params.TrainType != "" {
		query = query.Where("tt.train_type = ?", params.TrainType)
	}

	result := query.Order("s.time ASC").Find(&schedules)

	if result.Error != nil {
		return nil, result.Error
	}

	return schedules, nil
}
