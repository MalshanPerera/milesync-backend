package handlers

import (
	"jira-for-peasents/common"
	"jira-for-peasents/requests"
	"jira-for-peasents/responses"
	"jira-for-peasents/services"
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

	handler.POST("/register", h.registerUser)
}

func (h *AuthHandler) registerUser(c echo.Context) error {
	u := new(requests.RegisterUserRequest)
	ctx := c.Request().Context()
	if err := c.Bind(u); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(u); err != nil {
		return err
	}

	_, err := h.userService.GetUserFromEmail(ctx, u.Email)

	if err != nil && err != common.NoResults {
		return err
	}

	if err == nil {
		return common.BadRequest("User already exists")
	}

	newUser, err := h.userService.CreateUser(ctx, services.CreateUserParams{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, responses.NewRegisterUserResponse(
		newUser.ID,
		newUser.Name,
		newUser.Email,
	))

}
