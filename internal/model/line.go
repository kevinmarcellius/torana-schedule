package model

type StationInfo struct {
	Name     string `json:"name"`
	Distance int    `json:"distance"`
}

type LineDetail struct {
	Name     string        `json:"name"`
	Stations []StationInfo `json:"stations"`
}

type LinesResponse struct {
	Lines []LineDetail `json:"lines"`
}

// TrainTrip is used for internal data transfer from repository to service.
type TrainTrip struct {
	Line     string
	Station  string
	Distance int
	TrainType string
}
