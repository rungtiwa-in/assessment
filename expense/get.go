package expense

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
	"github.com/lib/pq"
)

func GetExpenseHandler(c echo.Context) error {
	id := c.Param("id")
	stmt, err := db.Prepare("SELECT id, title, amount, note, tags FROM expenses WHERE id = $1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare query user statement:" + err.Error()})
	}

	row := stmt.QueryRow(id)
	ex := Expense{}
	err = row.Scan(&ex.ID, &ex.Title, &ex.Amount, &ex.Note, pq.Array(&ex.Tags))
	switch err {
	case sql.ErrNoRows:
		return c.JSON(http.StatusNotFound, Err{Message: "expenses not found"})
	case nil:
		return c.JSON(http.StatusOK, ex)
	default:
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't scan expenses:" + err.Error()})
	}
}
