package main

import (
	"jira-for-peasents/server"
	"jira-for-peasents/utils"
)

func main() {
	httpServer := server.NewServer()
	if httpServer != nil {
		utils.InitLogger()
		logger := utils.GetLogger()
		httpServer.SetupLogger()
		httpServer.SetupCors()
		httpServer.SetupValidator()
		server.ConfigureRoutes(httpServer.Echo, logger)
		httpServer.Start()
	}
}
