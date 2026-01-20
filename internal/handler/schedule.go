package handler

import (
	"net/http"
	"time"

	"github.com/kevinmarcellius/torana-schedule/internal/model"
	"github.com/kevinmarcellius/torana-schedule/internal/service"
	"github.com/labstack/echo/v4"
)

type ScheduleHandler struct {
	scheduleSvc *service.ScheduleService
}

func NewScheduleHandler(scheduleSvc *service.ScheduleService) *ScheduleHandler {
	return &ScheduleHandler{scheduleSvc: scheduleSvc}
}

const timeLayout = "15:04:05"

// GetAllSchedules handles the GET /api/v1/schedules endpoint.
func (h *ScheduleHandler) GetAllSchedules(c echo.Context) error {
	groupedSchedules, err := h.scheduleSvc.GetGroupedSchedules()
	if err != nil {
		
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve schedules"})
	}
	return c.JSON(http.StatusOK, groupedSchedules)
}

// SearchSchedules handles the GET /api/v1/schedules/search endpoint.
func (h *ScheduleHandler) SearchSchedules(c echo.Context) error {
	if len(c.QueryParams()) == 0 {
		return h.GetAllSchedules(c)
	}

	params := &model.ScheduleSearchParams{
		Station:   c.QueryParam("station"),
		Line:      c.QueryParam("line"),
		TrainType: c.QueryParam("trainType"),
	}

	startTimeStr := c.QueryParam("startTime")
	if startTimeStr != "" {
		startTime, err := time.Parse(timeLayout, startTimeStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid startTime format. Use HH:MM:SS"})
		}
		params.StartTime = startTime
	}

	endTimeStr := c.QueryParam("endTime")
	if endTimeStr != "" {
		endTime, err := time.Parse(timeLayout, endTimeStr)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid endTime format. Use HH:MM:SS"})
		}
		params.EndTime = endTime
	}

	schedules, err := h.scheduleSvc.SearchSchedules(params)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve schedules"})
	}

	return c.JSON(http.StatusOK, schedules)
}
