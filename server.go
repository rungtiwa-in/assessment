package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo"
	"github.com/rungtiwa-in/assessment/expense"
)

func main() {
	expense.InitDB()

	e := echo.New()

	e.GET("/health", healthHandler)
	e.POST("/expenses", expense.CreateExpenseHandler)

	log.Println("Server started at :2565")
	log.Fatal(e.Start(":2565"))
}

func healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}
