package middlewares

import (
	errpkg "jira-for-peasants/errors"
	"jira-for-peasants/services"
	"jira-for-peasants/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CheckOrganization(organizationService *services.OrganizationService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			organizationId := c.Request().Header.Get("organization_id")
			userId := utils.GetUser(c)
			if userId == "" {
				return errpkg.NewApiError(http.StatusUnauthorized, "Unauthorized")
			}

			if organizationId == "" {
				return errpkg.BadRequest("Organization id is missing")
			}

			_, e := organizationService.GetOrganizationForUser(c.Request().Context(), userId, organizationId)
			if e != nil {
				return errpkg.NewApiError(http.StatusUnauthorized, "Unauthorized")
			}

			utils.SetOrganization(c, organizationId)

			return next(c)
		}
	}
}
