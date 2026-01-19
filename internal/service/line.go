package service

import (
	"sort"

	"github.com/kevinmarcellius/torana-schedule/internal/model"
	"github.com/kevinmarcellius/torana-schedule/internal/repository"
)

type LineService struct {
	lineRepo *repository.LineRepository
}

func NewLineService(lineRepo *repository.LineRepository) *LineService {
	return &LineService{lineRepo: lineRepo}
}

func (s *LineService) GetLinesWithStations() (*model.LinesResponse, error) {
	trips, err := s.lineRepo.GetLinesAndStations()
	if err != nil {
		return nil, err
	}

	linesMap := make(map[string]map[string]int)
	for _, trip := range trips {
		if _, ok := linesMap[trip.Line]; !ok {
			linesMap[trip.Line] = make(map[string]int)
		}
		// Avoid duplicates of stations, even if they have different train types
		linesMap[trip.Line][trip.Station] = trip.Distance
	}

	var lineDetails []model.LineDetail
	for lineName, stationsMap := range linesMap {
		var stations []model.StationInfo
		for stationName, distance := range stationsMap {
			stations = append(stations, model.StationInfo{
				Name:     stationName,
				Distance: distance,
			})
		}

		// Sort stations by distance
		sort.Slice(stations, func(i, j int) bool {
			return stations[i].Distance < stations[j].Distance
		})

		lineDetails = append(lineDetails, model.LineDetail{
			Name:     lineName,
			Stations: stations,
		})
	}


	return &model.LinesResponse{Lines: lineDetails}, nil
}
