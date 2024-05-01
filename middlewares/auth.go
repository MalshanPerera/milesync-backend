package middlewares

import (
	"jira-for-peasants/errors"
	"jira-for-peasants/utils"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type key string

const (
	UserIDKey key = "user_id"
)

func IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Check the Authorization header
		headerStr := c.Request().Header.Get("Authorization")

		// If the token is not in the Authorization header, check the cookies
		if headerStr == "" {
			cookie, err := c.Cookie("token")
			if err != nil {
				return errors.ApiError{
					Code:    http.StatusUnauthorized,
					Message: "Unauthorized",
				}
			}
			headerStr = cookie.Value
		}

		// Parse and validate the token
		tokenString := strings.TrimPrefix(headerStr, "Bearer ")
		token, err := utils.VerifyToken(tokenString)
		if err != nil {
			return errors.BadRequest("Invalid token")
		}

		// Extract the user information from the token
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return errors.BadRequest("Invalid token claims")
		}

		userId, ok := claims["user_id"].(string)
		if !ok {
			return errors.BadRequest("Invalid user id")
		}

		// Add the user information to the context
		utils.SetUser(c, userId)

		// If the user is authenticated, call the next handler with the new context
		return next(c)
	}
}
