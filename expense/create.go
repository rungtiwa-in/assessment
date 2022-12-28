package expense

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/lib/pq"
)

func CreateExpenseHandler(c echo.Context) error {
	ex := Expense{}
	err := c.Bind(&ex)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: err.Error()})
	}

	row := db.QueryRow("INSERT INTO expenses (title, amount, note, tags) values ($1, $2, $3, $4) RETURNING id, title, amount, note, tags;",
		ex.Title, ex.Amount, ex.Note, pq.Array(&ex.Tags))
	err = row.Scan(&ex.ID, &ex.Title, &ex.Amount, &ex.Note, pq.Array(&ex.Tags))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, ex)
}
