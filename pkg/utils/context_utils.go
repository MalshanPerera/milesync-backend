package utils

import "github.com/labstack/echo/v4"

func GetUser(ctx echo.Context) string {
	return ctx.Get("user_id").(string)
}

func SetUser(ctx echo.Context, user string) {
	ctx.Set("user_id", user)
}

func GetOrganization(ctx echo.Context) string {
	return ctx.Get("organization_id").(string)
}

func SetOrganization(ctx echo.Context, organization string) {
	ctx.Set("organization_id", organization)
}
