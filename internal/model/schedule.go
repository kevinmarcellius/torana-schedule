package model

import "time"

// Schedule represents a single train schedule entry in the search results.
type Schedule struct {
	Line      string `json:"line"`
	Station   string `json:"station"`
	TrainType string `json:"train_type"`
	Time      string `json:"time"` // Using string for HH:MM:SS format
}

// ScheduleSearchParams holds the parsed and validated search parameters.
// Zero values (e.g., empty string, zero time) mean the parameter is not set.
type ScheduleSearchParams struct {
	Station   string
	StartTime time.Time
	EndTime   time.Time
	Line      string
	TrainType string
}

// --- Models for Grouped Schedules Response ---

// ScheduleInfo is a compact schedule representation for the grouped view.
type ScheduleInfo struct {
	Line      string `json:"line"`
	TrainType string `json:"train_type"`
	Time      string `json:"time"`
}

// StationSchedules represents a station and its list of schedules.
type StationSchedules struct {
	Name      string         `json:"name"`
	Schedules []ScheduleInfo `json:"schedules"`
}

// GroupedSchedulesResponse is the top-level response for the grouped schedules view.
type GroupedSchedulesResponse struct {
	Stations []StationSchedules `json:"stations"`
}
