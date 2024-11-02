package handlers

import (
	"jira-for-peasants/internal/api/requests"
	"jira-for-peasants/internal/api/responses"
	"jira-for-peasants/internal/services"
	"jira-for-peasants/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type LabelHandler struct {
	labelService *services.LabelService
}

func NewLabelHandler(labelService *services.LabelService) LabelHandler {
	return LabelHandler{
		labelService: labelService,
	}
}

func (h *LabelHandler) RegisterRoutes(handler *echo.Group) {
	handler.POST("", h.handleCreateLabel)
	handler.PATCH("", h.handleUpdateLabel)
	handler.DELETE("", h.handleDeleteLabel)
	handler.GET("", h.handleGetLabels)
}

func (h *LabelHandler) handleCreateLabel(c echo.Context) error {
	u := new(requests.CreateLabelRequest)
	ctx := c.Request().Context()

	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(u); err != nil {
		return err
	}

	label, err := h.labelService.CreateLabel(ctx, services.CreateLabelParams{
		Name:           u.Name,
		Color:          u.Color,
		ProjectID:      u.ProjectId,
		OrganizationID: utils.GetOrganization(c),
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, responses.LabelResponse{
		ID:        label.ID,
		Name:      label.Name,
		Color:     label.Color,
		ProjectID: label.ProjectID,
	})
}

func (h *LabelHandler) handleUpdateLabel(c echo.Context) error {
	u := new(requests.UpdateLabelRequest)
	ctx := c.Request().Context()

	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(u); err != nil {
		return err
	}

	label, err := h.labelService.UpdateLabel(ctx, services.UpdateLabelParams{
		ID:    u.ID,
		Name:  u.Name,
		Color: u.Color,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, responses.LabelResponse{
		ID:        label.ID,
		Name:      label.Name,
		Color:     label.Color,
		ProjectID: label.ProjectID,
	})
}

func (h *LabelHandler) handleDeleteLabel(c echo.Context) error {
	return nil
}

func (h *LabelHandler) handleGetLabels(c echo.Context) error {
	return nil
}
