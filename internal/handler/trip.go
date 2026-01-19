package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/kevinmarcellius/torana-schedule/internal/service"
)

type TripHandler struct {
	service *service.TripService
}

func NewTripHandler(service *service.TripService) *TripHandler {
	return &TripHandler{service: service}
}

func (h *TripHandler) GetTripDetails(c echo.Context) error {
	source := c.QueryParam("source")
	destination := c.QueryParam("destination")
	trainType := c.QueryParam("trainType")

	if source == "" || destination == "" || trainType == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "source, destination, and trainType query parameters are required"})
	}

	distance, err := h.service.GetTripDistance(source, destination, trainType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"distance": distance})
}
