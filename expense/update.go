package expense

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/lib/pq"
)

func UpdateExpenseHandler(c echo.Context) error {
	id := c.Param("id")
	stmt, err := db.Prepare("UPDATE expenses SET title=$2, amount=$3, note=$4, tags=$5 WHERE id = $1")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't prepare query expenses statement:" + err.Error()})
	}

	ex := Expense{}
	err = c.Bind(&ex)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "can't bind data expenses statement:" + err.Error()})
	}

	row, err := stmt.Exec(id, &ex.Title, &ex.Amount, &ex.Note, pq.Array(&ex.Tags))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't update expenses statement:" + err.Error()})
	}

	if _, err := row.RowsAffected(); err != nil {
		return c.JSON(http.StatusInternalServerError, Err{Message: "can't get expenses row statement:" + err.Error()})
	}

	cid, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Err{Message: "can't convert expenses id:" + err.Error()})
	}

	ex.ID = cid
	return c.JSON(http.StatusOK, ex)
}
