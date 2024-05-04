package handlers

import (
	errpkg "jira-for-peasants/errors"
	"jira-for-peasants/requests"
	"jira-for-peasants/responses"
	"jira-for-peasants/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	userService *services.UserService
}

func NewAuthHandler(userService *services.UserService) AuthHandler {
	return AuthHandler{
		userService: userService,
	}
}

func (h *AuthHandler) RegisterRoutes(handler *echo.Group) {

	handler.POST("/register", h.handleRegisterUser)
	handler.POST("/login", h.handleLoginUser)
}

func (h *AuthHandler) handleRegisterUser(c echo.Context) error {
	u := new(requests.RegisterUserRequest)
	ctx := c.Request().Context()
	if err := c.Bind(u); err != nil {
		return errpkg.BadRequest(err.Error())
	}
	if err := c.Validate(u); err != nil {
		return errpkg.UnprocessableEntity(err.Error())
	}

	newUser, newSession, err := h.userService.CreateUser(ctx, services.CreateUserParams{
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
		Password:  u.Password,
	})

	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, responses.NewAuthResponse(
		newUser.ID,
		newUser.FirstName,
		newUser.LastName,
		newUser.Email,
		newSession.AccessToken,
		newSession.RefreshToken,
		newSession.ExpiresAt,
	))
}

func (h *AuthHandler) handleLoginUser(c echo.Context) error {
	u := new(requests.LoginUserRequest)
	ctx := c.Request().Context()
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(u); err != nil {
		return err
	}

	existingUser, newSession, err := h.userService.LoginUser(ctx, services.LoginUserParams{
		Email:    u.Email,
		Password: u.Password,
	})
	if err != nil {
		return errpkg.BadRequest(errpkg.InvalidPassword)
	}

	return c.JSON(http.StatusOK, responses.NewAuthResponse(
		existingUser.ID,
		existingUser.FirstName,
		existingUser.LastName,
		existingUser.Email,
		newSession.AccessToken,
		newSession.RefreshToken,
		newSession.ExpiresAt,
	))
}
