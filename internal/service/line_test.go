package service

import (
	"errors"
	"sort"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kevinmarcellius/torana-schedule/internal/model"
	mock_repository "github.com/kevinmarcellius/torana-schedule/internal/repository/mocks"
	"github.com/stretchr/testify/assert"
)

func TestLineService_GetLinesWithStations(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockLineRepo := mock_repository.NewMockLineRepository(ctrl)

	testCases := []struct {
		name          string
		mockTrips     []model.TrainTrip
		mockErr       error
		expectedLines *model.LinesResponse
		expectedErr   error
	}{
		{
			name: "Success",
			mockTrips: []model.TrainTrip{
				{Line: "blue", Station: "A", Distance: 0},
				{Line: "blue", Station: "B", Distance: 10},
				{Line: "red", Station: "C", Distance: 0},
				{Line: "red", Station: "D", Distance: 12},
			},
			mockErr: nil,
			expectedLines: &model.LinesResponse{
				Lines: []model.LineDetail{
					{
						Name: "blue",
						Stations: []model.StationInfo{
							{Name: "A", Distance: 0},
							{Name: "B", Distance: 10},
						},
					},
					{
						Name: "red",
						Stations: []model.StationInfo{
							{Name: "C", Distance: 0},
							{Name: "D", Distance: 12},
						},
					},
				},
			},
			expectedErr: nil,
		},
		{
			name:          "Repository error",
			mockTrips:     nil,
			mockErr:       errors.New("repository error"),
			expectedLines: nil,
			expectedErr:   errors.New("repository error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockLineRepo.EXPECT().GetLinesAndStations().Return(tc.mockTrips, tc.mockErr)

			s := NewLineService(mockLineRepo)
			lines, err := s.GetLinesWithStations()

			assert.Equal(t, tc.expectedErr, err)

			if tc.expectedLines != nil {
				assert.NotNil(t, lines)

				// Sort both expected and actual slices for consistent comparison
				sortLinesResponse(tc.expectedLines)
				sortLinesResponse(lines)

				assert.Equal(t, *tc.expectedLines, *lines)
			} else {
				assert.Nil(t, lines)
			}
		})
	}
}

// Helper function to sort the LinesResponse for consistent comparison in tests
func sortLinesResponse(resp *model.LinesResponse) {
	if resp == nil {
		return
	}
	for i := range resp.Lines {
		sort.Slice(resp.Lines[i].Stations, func(j, k int) bool {
			return resp.Lines[i].Stations[j].Distance < resp.Lines[i].Stations[k].Distance
		})
	}
	sort.Slice(resp.Lines, func(i, j int) bool {
		return resp.Lines[i].Name < resp.Lines[j].Name
	})
}
