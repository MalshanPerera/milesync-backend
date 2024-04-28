package middlewares

import (
	"jira-for-peasents/common"

	"github.com/labstack/echo/v4"
)

func LoggingMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// log the request
		common.Logger.LogInfo().Fields(map[string]interface{}{
			"method": c.Request().Method,
			"uri":    c.Request().URL.Path,
			"query":  c.Request().URL.RawQuery,
		}).Msg("Request")

		// call the next middleware/handler
		err := next(c)
		if err != nil {

			if e, ok := err.(common.ApiError); ok {
				common.Logger.LogInfo().Fields(map[string]interface{}{
					"info": e.Error(),
				}).Msg("Response")
				return e
			}

			common.Logger.LogError().Fields(map[string]interface{}{
				"error": err.Error(),
			}).Msg("Response")
			return err
		}

		return nil
	}
}
