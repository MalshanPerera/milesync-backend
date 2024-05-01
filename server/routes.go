package server

import (
	"jira-for-peasents/common"
	"jira-for-peasents/handlers"
	"jira-for-peasents/middlewares"
	"jira-for-peasents/services"
)

func ConfigureRoutes(
	s *Server,
) {
	apiV1 := s.Echo.Group("/api/v1")

	// Registering services
	userService := services.NewUserService(s.DB)
	projectService := services.NewProjectService(s.DB)
	organizationService := services.NewOrganizationService(s.DB)

	// Registering handlers
	authHandler := handlers.NewAuthHandler(userService)
	projectHandler := handlers.NewProjectHandler(projectService)
	organizationHandler := handlers.NewOrganizationHandler(organizationService)

	// Registering routes
	authGroup := apiV1.Group("/auth")
	authHandler.RegisterRoutes(authGroup)

	protectedRoutes := apiV1.Group("")
	protectedRoutes.Use(middlewares.IsAuthenticated)

	projectGroup := protectedRoutes.Group("/projects")
	projectHandler.RegisterRoutes(projectGroup)

	organizationGroup := protectedRoutes.Group("/organizations")
	organizationHandler.RegisterRoutes(organizationGroup)

	routeList := s.Echo.Routes()
	for _, route := range routeList {
		if route.Method != "echo_route_not_found" {
			common.Logger.LogInfo().Msg("Route Registered: " + route.Path + " " + route.Method)
		}
	}
}
