package handler

import (
	"net/http"

	"github.com/kevinmarcellius/torana-schedule/internal/service"
	"github.com/labstack/echo/v4"
)

type StationHandler struct {
	stationSvc *service.StationService
}

func NewStationHandler(stationSvc *service.StationService) *StationHandler {
	return &StationHandler{stationSvc: stationSvc}
}

func (h *StationHandler) GetStationsByLine(c echo.Context) error {
	lineName := c.Param("lineName")
	if lineName == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Line name is required"})
	}

	response, err := h.stationSvc.GetLineWithStations(lineName)
	if err != nil {
		// In a real app, you'd want to log this error
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve stations"})
	}

	if response == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "No stations found for this line"})
	}

	return c.JSON(http.StatusOK, response)
}
