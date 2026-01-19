package model

// StationWithSchedule represents a station with its scheduled arrival/departure time.
type StationWithSchedule struct {
	Name          string `json:"name"`
	Distance      int    `json:"distance"`
	ScheduledTime string `json:"scheduled_time"`
}

// LineWithStations is the response model for the GET /api/v1/lines/{lineName} endpoint.
type LineWithStations struct {
	Name     string                `json:"name"`
	Stations []StationWithSchedule `json:"stations"`
}

// StationScheduleData is an internal struct to transfer data from the repository.
type StationScheduleData struct {
	StationName   string
	Distance      int
	ScheduledTime string
}
