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

	sessionService := services.NewSessionService(sessionRepository)
	projectService := services.NewProjectService(projectsRepository)
	organizationService := services.NewOrganizationService(organizationRepository)

	// Registering handlers
	authHandler := handlers.NewAuthHandler(userService)
	organizationHandler := handlers.NewOrganizationHandler(organizationService)
	projectHandler := handlers.NewProjectHandler(projectService)

	// Registering routes
	authGroup := apiV1.Group("/auth")
	authHandler.RegisterRoutes(authGroup)

	protectedRoutes := apiV1.Group("")
	protectedRoutes.Use(middlewares.IsAuthenticated(sessionService))

	organizationGroup := protectedRoutes.Group("/organizations")
	organizationHandler.RegisterRoutes(organizationGroup)

	projectGroup := protectedRoutes.Group("/projects")
	projectHandler.RegisterRoutes(projectGroup)

	routeList := s.Echo.Routes()
	for _, route := range routeList {
		if route.Method != "echo_route_not_found" {
			common.Logger.LogInfo().Msg("Route Registered: " + route.Path + " " + route.Method)
		}
	}
}
