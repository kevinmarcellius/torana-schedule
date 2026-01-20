package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/kevinmarcellius/torana-schedule/config"
	"github.com/kevinmarcellius/torana-schedule/internal/handler"
	"github.com/kevinmarcellius/torana-schedule/internal/repository"
	"github.com/kevinmarcellius/torana-schedule/internal/service"
)

func main() {
	fmt.Println("Hello, world!")

	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	hello := cfg.Postgres.Host

	output := "Hello " + hello

	log.Println(output)

	db, err := config.ConnectPostgres(cfg.Postgres)
	if err != nil {
		log.Fatalf("Failed to connect to Postgres: %v", err)
	}

	err = config.DBHealthCheck(db)
	if err != nil {
		log.Fatalf("Database health check failed: %v", err)
	}
	log.Println("Database connection is healthy.")

	healthHandler := handler.NewHealthHandler(db)

	// Line endpoint wiring
	lineRepo := repository.NewLineRepository(db)
	lineService := service.NewLineService(lineRepo)
	lineHandler := handler.NewLineHandler(lineService)

	// Station endpoint wiring
	stationRepo := repository.NewStationRepository(db)
	stationService := service.NewStationService(stationRepo)
	stationHandler := handler.NewStationHandler(stationService)

	// Schedule endpoint wiring
	scheduleRepo := repository.NewScheduleRepository(db)
	scheduleService := service.NewScheduleService(scheduleRepo)
	scheduleHandler := handler.NewScheduleHandler(scheduleService)

	// Trip endpoint wiring
	tripRepo := repository.NewTripRepository(db)
	tripService := service.NewTripService(tripRepo)
	tripHandler := handler.NewTripHandler(tripService)

	e := echo.New()
	e.Use(middleware.Logger()) // Add this line to enable the logger middleware
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, output)
	})

	v1 := e.Group("/api/v1")
	v1.GET("/health/ready", healthHandler.ReadinessCheck)
	v1.GET("/health/live", healthHandler.LivenessCheck)
	v1.GET("/lines", lineHandler.GetLines)
	v1.GET("/lines/:lineName", stationHandler.GetStationsByLine)
	v1.GET("/schedules/search", scheduleHandler.SearchSchedules)
	v1.GET("/trips/details/search", tripHandler.GetTripDetails)

	port := fmt.Sprintf(":%d", cfg.Port)
	log.Printf("Starting server on port %s\n", port)
	e.Logger.Fatal(e.Start(port))
}
