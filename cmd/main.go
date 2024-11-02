package main

import (
	"jira-for-peasants/internal/middlewares"
	"jira-for-peasants/internal/server"
)

func main() {
	httpServer := server.NewServer()
	if httpServer != nil {
		httpServer.Echo.Use(middlewares.LoggingMiddleware)
		httpServer.SetupCors()
		httpServer.SetupErrorHandler()
		httpServer.SetupValidator()
		server.ConfigureRoutes(httpServer)
		httpServer.Start()
	}
}
