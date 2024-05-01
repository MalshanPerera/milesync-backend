package handlers

import (
	"jira-for-peasents/common"
	"jira-for-peasents/requests"
	"jira-for-peasents/responses"
	"jira-for-peasents/services"
	"jira-for-peasents/utils"
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

func (h *OrganizationHandler) RegisterRoutes(handler *echo.Group) {
	handler.POST("", h.createOrganization)
	handler.GET("/available", h.getHandleAvailable)
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
		return common.BadRequest("ORGANIZATION_NAME_ALREADY_USED")
	}

	newOrg, err := h.organizationService.CreateOrganization(ctx, services.CreateOrganizationParams{
		UserId: utils.GetUser(c),
		Name:   u.Name,
	})

	if err != nil {
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
		return common.BadRequest("ORGANIZATION_NAME_REQUIRED")
	}

	status, e := h.organizationService.GetOrganizationSlugUsed(c.Request().Context(), name)

	if e != nil {
		return e
	}

	return c.JSON(http.StatusOK, !status)
}
