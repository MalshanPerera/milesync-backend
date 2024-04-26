package server

import (
	"fmt"
	"jira-for-peasents/config"

	"github.com/labstack/echo/v4"
)

type Server struct {
	Echo   *echo.Echo
	Config *config.Config
}

func NewServer() *Server {
	config := config.NewConfig()
	err := config.Validate()
	if err == nil {
		return &Server{
			Echo:   echo.New(),
			Config: config,
		}
	}
	fmt.Printf("Invalid config %s \n", err.Error())

	return nil
}

func (s *Server) Start() {
	s.Echo.Logger.Fatal(s.Echo.Start(":" + s.Config.Port))
}
