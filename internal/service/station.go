package service

import (
	"github.com/kevinmarcellius/torana-schedule/internal/model"
	"github.com/kevinmarcellius/torana-schedule/internal/repository"
)

type StationService struct {
	stationRepo repository.StationRepository
}

func NewStationService(stationRepo repository.StationRepository) *StationService {
	return &StationService{stationRepo: stationRepo}
}

func (s *StationService) GetLineWithStations(lineName string) (*model.LineWithStations, error) {
	data, err := s.stationRepo.GetStationsAndSchedulesByLine(lineName)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		// Return nil to indicate that the line was not found.
		return nil, nil
	}

	stations := make([]model.StationWithSchedule, len(data))
	for i, item := range data {
		stations[i] = model.StationWithSchedule{
			Name:          item.StationName,
			Distance:      item.Distance,
			ScheduledTime: item.ScheduledTime,
		}
	}

	return &model.LineWithStations{
		Name:     lineName,
		Stations: stations,
	}, nil
}
