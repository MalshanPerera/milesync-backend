package server

import (
	"context"
	"errors"
	"fmt"
	"jira-for-peasants/config"
	datastore "jira-for-peasants/db"
	errpkg "jira-for-peasants/errors"
	"jira-for-peasants/utils"
	"net/http"
	"os"
	"os/signal"
	"syscall"
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
		if validatorErr, ok := err.(validator.ValidationErrors); ok {
			// return error from first failure, may want to return list with all errors
			errorMap := make(map[string][]string)
			for _, err := range validatorErr {
				fieldName := err.Field()
				errorMap[fieldName] = append(errorMap[fieldName], err.Tag())
			}
			return errpkg.UnprocessableEntity(errorMap)
		}
		return err
	}
	return nil
}

func NewServer() *Server {
	cfg := config.NewConfig()
	err := cfg.Validate()
	if err == nil {
		utils.InitJwt()
		db := datastore.NewDB(cfg.DB.User, cfg.DB.Password, cfg.DB.Host, cfg.DB.Port, cfg.DB.DatabaseName)
		return &Server{
			Echo:   echo.New(),
			Config: cfg,
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
		if myErr, ok := err.(errpkg.ApiError); ok {
			if err := c.JSON(myErr.Code, myErr); err != nil {
				s.Echo.Logger.Error(err)
			}
			return
		}

		// TODO: handle other error types
		// use default error handler functionality
		s.Echo.DefaultHTTPErrorHandler(err, c)
	}
}

func (s *Server) SetupCors() {
	s.Echo.Use(middleware.CORS())
}

func (s *Server) Start() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	// Start server
	go func() {
		if err := s.Echo.Start(":" + s.Config.Port); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.Echo.Logger.Error("error starting server:", err) // Consider using Error level here
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
			s.Echo.Logger.Fatal("error during shutdown:", err)
		}
	case <-ctxShutdown.Done():
		s.Echo.Logger.Fatal("server shutdown timed out")
	}
}
