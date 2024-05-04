package server

import (
	"context"
	"fmt"
	"jira-for-peasants/config"
	datastore "jira-for-peasants/db"
	err_pkg "jira-for-peasants/errors"
	"jira-for-peasants/utils"
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
	DB     *datastore.DB
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
		utils.InitJwt()
		db := datastore.NewDB(config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.DatabaseName)
		return &Server{
			Echo:   echo.New(),
			Config: config,
			DB:     db,
		}
	}
	fmt.Printf("Invalid config %s \n", err.Error())

	return nil
}

func (s *Server) SetupValidator() {
	s.Echo.Validator = &CustomValidator{validator: validator.New()}
}

func (s *Server) SetupErrorHandler() {
	s.Echo.HTTPErrorHandler = func(err error, c echo.Context) {
		if c.Response().Committed {
			return
		}

		// check if error is known type to be handled differently
		if myErr, ok := err.(err_pkg.ApiError); ok {
			if err := c.JSON(myErr.Code, map[string]string{
				"code":    fmt.Sprintf("%d", myErr.Code),
				"message": myErr.Message,
			}); err != nil {
				s.Echo.Logger.Error(err)
			}
			return
		}

		//TODO: handle other error types
		// use default error handler functionality
		s.Echo.DefaultHTTPErrorHandler(err, c)
	}
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
