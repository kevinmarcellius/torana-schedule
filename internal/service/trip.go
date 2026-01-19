package service

import (
	"container/heap"
	"fmt"
	"math"
	"sort"

	"github.com/kevinmarcellius/torana-schedule/internal/model"
	"github.com/kevinmarcellius/torana-schedule/internal/repository"
)

// TripService handles the business logic for trips.
type TripService struct {
	tripRepo *repository.TripRepository
}

// NewTripService creates a new TripService.
func NewTripService(tripRepo *repository.TripRepository) *TripService {
	return &TripService{tripRepo: tripRepo}
}

// GetTripDistance calculates the shortest distance between two stations.
func (s *TripService) GetTripDistance(sourceStation, destinationStation, trainType string) (int, error) {
	graph, err := s.buildGraph(trainType)
	if err != nil {
		return 0, err
	}

	return s.dijkstra(graph, sourceStation, destinationStation)
}

// buildGraph constructs a graph from the trip data in the database.
func (s *TripService) buildGraph(trainType string) (map[string]map[string]int, error) {
	trips, err := s.tripRepo.GetAllTrips()
	if err != nil {
		return nil, err
	}

	// Filter trips by trainType
	var filteredTrips []model.TrainTrip
	for _, trip := range trips {
		if trip.TrainType == trainType {
			filteredTrips = append(filteredTrips, trip)
		}
	}

	// Group trips by line
	tripsByLine := make(map[string][]model.TrainTrip)
	for _, trip := range filteredTrips {
		tripsByLine[trip.Line] = append(tripsByLine[trip.Line], trip)
	}

	graph := make(map[string]map[string]int)

	// Process each line to build graph edges
	for _, lineTrips := range tripsByLine {
		// Sort stations by distance to ensure they are in order
		sort.Slice(lineTrips, func(i, j int) bool {
			return lineTrips[i].Distance < lineTrips[j].Distance
		})

		for i := 0; i < len(lineTrips)-1; i++ {
			stationA := lineTrips[i]
			stationB := lineTrips[i+1]

			if graph[stationA.Station] == nil {
				graph[stationA.Station] = make(map[string]int)
			}
			if graph[stationB.Station] == nil {
				graph[stationB.Station] = make(map[string]int)
			}

			distance := int(math.Abs(float64(stationB.Distance - stationA.Distance)))
			graph[stationA.Station][stationB.Station] = distance
			graph[stationB.Station][stationA.Station] = distance
		}
	}

	return graph, nil
}

// dijkstra finds the shortest path between two nodes in a graph.
func (s *TripService) dijkstra(graph map[string]map[string]int, start, end string) (int, error) {
	distances := make(map[string]int)
	for station := range graph {
		distances[station] = math.MaxInt32
	}
	distances[start] = 0

	pq := make(PriorityQueue, 0)
	heap.Push(&pq, &Item{value: start, priority: 0})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		station := item.value

		if station == end {
			return distances[end], nil
		}

		for neighbor, weight := range graph[station] {
			distance := distances[station] + weight
			if distance < distances[neighbor] {
				distances[neighbor] = distance
				heap.Push(&pq, &Item{value: neighbor, priority: distance})
			}
		}
	}

	return 0, fmt.Errorf("no path found from %s to %s", start, end)
}

// PriorityQueue implementation for Dijkstra's
type Item struct {
	value    string
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}
