package main

import (
	"jira-for-peasents/common"
	"jira-for-peasents/middlewares"
	"jira-for-peasents/server"
)

func main() {
	httpServer := server.NewServer()
	if httpServer != nil {
		common.NewLogger()
		httpServer.Echo.Use(middlewares.LoggingMiddleware)
		httpServer.SetupCors()
		httpServer.SetupErrorHandler()
		httpServer.SetupValidator()
		server.ConfigureRoutes(httpServer)
		httpServer.Start()
	}
}
