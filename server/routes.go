package server

import (
	"jira-for-peasents/common"
	"jira-for-peasents/handlers"
	"jira-for-peasents/services"
)

func ConfigureRoutes(
	s *Server,
) {
	apiV1 := s.Echo.Group("/api/v1")
	userService := services.NewUserService(s.DB)

	authHandler := handlers.NewAuthHandler(userService)

	authGroup := apiV1.Group("/auth")
	authHandler.RegisterRoutes(authGroup)

	routeList := s.Echo.Routes()
	for _, route := range routeList {
		common.Logger.LogInfo().Msg("Route Registered: " + route.Path + " " + route.Method)
	}
}
