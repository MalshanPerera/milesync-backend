package server

import (
	"jira-for-peasants/common"
	"jira-for-peasants/handlers"
	"jira-for-peasants/middlewares"
	"jira-for-peasants/repositories"
	"jira-for-peasants/services"
)

func ConfigureRoutes(
	s *Server,
) {
	apiV1 := s.Echo.Group("/api/v1")

	userRepository := repositories.NewUserRepository(s.DB)
	sessionRepository := repositories.NewSessionRepository(s.DB)
	organizationRepository := repositories.NewOrganizationRepository(s.DB)
	projectsRepository := repositories.NewProjectRepository(s.DB)

	// Registering services
	userService := services.NewUserService(
		userRepository,
		sessionRepository,
	)

	projectService := services.NewProjectService(projectsRepository)
	organizationService := services.NewOrganizationService(organizationRepository)

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
