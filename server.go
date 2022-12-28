package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/rungtiwa-in/assessment/expense"
)

func main() {
	expense.InitDB()

	e := echo.New()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := c.Request().Header.Get("Authorization")
			if auth != "November 10, 2009" {
				return c.JSON(http.StatusUnauthorized, "Unauthorized")
			}

			return next(c)
		}
	})

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", healthHandler)
	e.POST("/expenses", expense.CreateExpenseHandler)

	log.Println("Server started at :2565")
	log.Fatal(e.Start(":2565"))
}

func healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}
