package handlers

import (
	"errors"
	"jira-for-peasants/internal/api/requests"
	"jira-for-peasants/internal/api/responses"
	"jira-for-peasants/internal/services"
	errpkg "jira-for-peasants/pkg/errors"
	"jira-for-peasants/pkg/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type OrganizationHandler struct {
	organizationService *services.OrganizationService
}

func NewOrganizationHandler(service *services.OrganizationService) OrganizationHandler {
	return OrganizationHandler{
		organizationService: service,
	}
}

func (h *OrganizationHandler) RegisterRoutes(handler *echo.Group, authMiddleware echo.MiddlewareFunc) {
	handler.POST("", h.createOrganization, authMiddleware)
	handler.GET("/available", h.getHandleAvailable, authMiddleware)
}

func (h *OrganizationHandler) createOrganization(c echo.Context) error {

	u := new(requests.CreateOrganizationRequest)
	ctx := c.Request().Context()
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(u); err != nil {
		return err
	}

	status, e := h.organizationService.GetOrganizationSlugUsed(ctx, u.Name)

	if e != nil {
		return e
	}

	if status {
		return errpkg.BadRequest(errpkg.OrganizationNameExists)
	}

	newOrg, err := h.organizationService.CreateOrganization(ctx, services.CreateOrganizationParams{
		UserId: utils.GetUser(c),
		Name:   u.Name,
	})

	if err != nil {
		var appErr errpkg.AppError
		if errors.As(err, &appErr) {
			return errpkg.BadRequest(appErr.Message)
		}
		return err
	}

	return c.JSON(http.StatusOK, responses.NewOrganizationResponse(
		newOrg.ID,
		newOrg.Name,
		newOrg.Slug,
		newOrg.UserID,
	))
}

func (h *OrganizationHandler) getHandleAvailable(c echo.Context) error {
	name := c.QueryParam("name")
	if name == "" {
		return errpkg.BadRequest(errpkg.OrganizationNameRequired)
	}

	status, e := h.organizationService.GetOrganizationSlugUsed(c.Request().Context(), name)

	if e != nil {
		return e
	}

	return c.JSON(http.StatusOK, !status)
}
