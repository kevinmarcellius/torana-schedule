package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/kevinmarcellius/torana-schedule/config"
	handler "github.com/kevinmarcellius/torana-schedule/internal/handler"
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

	e := echo.New()
	e.Use(middleware.Logger()) // Add this line to enable the logger middleware
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, output)
	})

	v1 := e.Group("/api/v1")
	v1.GET("/health/ready", healthHandler.ReadinessCheck)
	v1.GET("/health/live", healthHandler.LivenessCheck)

	port := fmt.Sprintf(":%d", cfg.Port)
	log.Printf("Starting server on port %s\n", port)
	e.Logger.Fatal(e.Start(port))
}
