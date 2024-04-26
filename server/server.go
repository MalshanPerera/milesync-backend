package server

import (
	"context"
	"fmt"
	"jira-for-peasents/config"
	"jira-for-peasents/utils"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	Echo   *echo.Echo
	Config *config.Config
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
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

func (s *Server) SetupValidator() {
	s.Echo.Validator = &CustomValidator{validator: validator.New()}
}

func (s *Server) SetupLogger() {
	logger := utils.GetZeroLogger()
	s.Echo.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info().
				Str("URI", v.URI).
				Int("status", v.Status).
				Msg("request")

			return nil
		},
	}))
}

func (s *Server) SetupCors() {
	s.Echo.Use(middleware.CORS())
}

func (s *Server) Start() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server
	go func() {
		if err := s.Echo.Start(":" + s.Config.Port); err != nil && err != http.ErrServerClosed {
			s.Echo.Logger.Fatal("shutting down the server", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctxShutdown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	shutdownError := make(chan error, 1)
	go func() {
		shutdownError <- s.Echo.Shutdown(ctxShutdown)
	}()

	select {
	case err := <-shutdownError:
		if err != nil {
			s.Echo.Logger.Fatal(err)
		}
	case <-ctxShutdown.Done():
		s.Echo.Logger.Fatal("shutdown timeout")
	}
}
