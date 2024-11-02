package handlers

import (
	"jira-for-peasants/internal/services"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	taskService *services.TaskService
}

func NewTaskHandler(taskService *services.TaskService) TaskHandler {
	return TaskHandler{
		taskService: taskService,
	}
}

func (h *TaskHandler) RegisterRoutes(handler *echo.Group) {
	handler.POST("", h.handleCreateTask)
	handler.PATCH("", h.handleUpdateTask)
	handler.DELETE("", h.handleDeleteTask)
	handler.GET("", h.handleGetTasks)
	handler.GET("/:id", h.handleGetTask)
}

func (h *TaskHandler) handleCreateTask(c echo.Context) error {
	return nil
}

func (h *TaskHandler) handleUpdateTask(c echo.Context) error {
	return nil
}

func (h *TaskHandler) handleDeleteTask(c echo.Context) error {
	return nil
}

func (h *TaskHandler) handleGetTasks(c echo.Context) error {
	return nil
}

func (h *TaskHandler) handleGetTask(c echo.Context) error {
	return nil
}
