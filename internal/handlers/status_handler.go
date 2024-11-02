package handlers

import (
	"jira-for-peasants/internal/services"

	"github.com/labstack/echo/v4"
)

type StatusHandler struct {
	statusService *services.StatusService
}

func NewStatusHandler(statusService *services.StatusService) StatusHandler {
	return StatusHandler{
		statusService: statusService,
	}
}

func (h *StatusHandler) RegisterRoutes(handler *echo.Group) {

}

func (h *StatusHandler) handleCreateStatus(c echo.Context) error {
	return nil
}

func (h *StatusHandler) handleUpdateStatus(c echo.Context) error {
	return nil
}

func (h *StatusHandler) handleDeleteStatus(c echo.Context) error {
	return nil
}

func (h *StatusHandler) handleGetStatuses(c echo.Context) error {
	return nil
}
