package handlers

import (
	"jira-for-peasents/requests"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct{}

func NewAuthHandler() AuthHandler {
	return AuthHandler{}
}

func (h *AuthHandler) RegisterRoutes(handler *echo.Group) {

	handler.POST("/register", h.registerUser)
}

func (h *AuthHandler) registerUser(c echo.Context) error {
	u := new(requests.RegisterUserRequest)
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)

}
