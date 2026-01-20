package handler

import (
	"net/http"

	"github.com/kevinmarcellius/torana-schedule/internal/service"
	"github.com/labstack/echo/v4"
)

type LineHandler struct {
	lineSvc *service.LineService
}

func NewLineHandler(lineSvc *service.LineService) *LineHandler {
	return &LineHandler{lineSvc: lineSvc}
}

func (h *LineHandler) GetLines(c echo.Context) error {
	linesResponse, err := h.lineSvc.GetLinesWithStations()
	if err != nil {
		
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve lines"})
	}
	return c.JSON(http.StatusOK, linesResponse)
}
