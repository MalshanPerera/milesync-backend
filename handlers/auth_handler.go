package handlers

import "github.com/labstack/echo/v4"

type AuthHandler struct{}

func NewAuthHandler() AuthHandler {
	return AuthHandler{}
}

func (h *AuthHandler) RegisterRoutes(handler *echo.Group) {
	handler.POST("/register", h.RegisterUser)
}

func (h *AuthHandler) RegisterUser(ctx echo.Context) error {
	return ctx.JSON(200, "User registered")
}
