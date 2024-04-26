package server

import (
	"jira-for-peasents/handlers"
	"jira-for-peasents/utils"

	"github.com/labstack/echo/v4"
)

func ConfigureRoutes(
	echo *echo.Echo,
	logger utils.Logger,
) {
	apiV1 := echo.Group("/api/v1")
	authHandler := handlers.NewAuthHandler()

	authGroup := apiV1.Group("/auth")
	authHandler.RegisterRoutes(authGroup)

	routeList := echo.Routes()
	for _, route := range routeList {
		logger.Info("Route Registered: " + route.Path + " " + route.Method)
	}
}
