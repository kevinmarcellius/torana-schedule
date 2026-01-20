package service

import (

	"github.com/kevinmarcellius/torana-schedule/internal/model"
	"github.com/kevinmarcellius/torana-schedule/internal/repository"
)

type ScheduleService struct {
	scheduleRepo repository.ScheduleRepository
}

func NewScheduleService(scheduleRepo repository.ScheduleRepository) *ScheduleService {
	return &ScheduleService{scheduleRepo: scheduleRepo}
}

func (s *ScheduleService) SearchSchedules(params *model.ScheduleSearchParams) ([]model.Schedule, error) {
	return s.scheduleRepo.SearchSchedules(params)
}

func (s *ScheduleService) GetGroupedSchedules() (*model.GroupedSchedulesResponse, error) {
	schedules, err := s.scheduleRepo.GetAllSchedules()
	if err != nil {
		return nil, err
	}

	grouped := make(map[string][]model.ScheduleInfo)
	for _, s := range schedules {
		info := model.ScheduleInfo{
			Line:      s.Line,
			TrainType: s.TrainType,
			Time:      s.Time,
		}
		grouped[s.Station] = append(grouped[s.Station], info)
	}

	var stationSchedules []model.StationSchedules
	for stationName, scheduleInfos := range grouped {
		stationSchedules = append(stationSchedules, model.StationSchedules{
			Name:      stationName,
			Schedules: scheduleInfos,
		})
	}


	return &model.GroupedSchedulesResponse{Stations: stationSchedules}, nil
}
