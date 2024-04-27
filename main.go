package main

import (
	"jira-for-peasents/common"
	"jira-for-peasents/server"
)

func main() {
	httpServer := server.NewServer()
	if httpServer != nil {
		common.NewLogger() // new
		httpServer.Echo.Use(common.LoggingMiddleware)
		httpServer.SetupCors()
		httpServer.SetupErrorHandler()
		httpServer.SetupValidator()
		server.ConfigureRoutes(httpServer)
		httpServer.Start()
	}
}
