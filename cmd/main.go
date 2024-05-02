package main

import (
	"jira-for-peasants/middlewares"
	"jira-for-peasants/server"
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
