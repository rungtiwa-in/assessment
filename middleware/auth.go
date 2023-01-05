package middleware

import (
	"net/http"

	"github.com/labstack/echo"
)

func Authorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		auth := c.Request().Header.Get("Authorization")
		if auth != "November 10, 2009" {
			return c.JSON(http.StatusUnauthorized, "Unauthorized")
		}

		return next(c)
	}
}
