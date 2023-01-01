package expense

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateExpense(t *testing.T) {
	body := bytes.NewBufferString(`{
		"title": "strawberry smoothie",
		"amount": 79,
		"note": "night market promotion discount 10 bath", 
		"tags": ["food", "beverage"]
	}`)
	var ex Expense

	res := request(http.MethodPost, uri("expenses"), body)
	err := res.Decode(&ex)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusCreated, res.StatusCode)
	assert.NotEqual(t, 0, ex.ID)
	assert.Equal(t, "strawberry smoothie", ex.Title)
	assert.Equal(t, float64(79), ex.Amount)
	assert.Equal(t, []string{"food", "beverage"}, ex.Tags)
}

func TestGetExpenseByID(t *testing.T) {
	e := seedExpense(t)

	var latest Expense
	res := request(http.MethodGet, uri("expenses", strconv.Itoa(e.ID)), nil)
	err := res.Decode(&latest)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, e.ID, latest.ID)
	assert.Equal(t, e.Title, latest.Title)
	assert.Equal(t, e.Amount, latest.Amount)
	assert.Equal(t, e.Note, latest.Note)
	assert.Equal(t, e.Tags, latest.Tags)
	assert.NotEmpty(t, latest.Title)
	assert.NotEmpty(t, latest.Amount)
	assert.NotEmpty(t, latest.Note)
	assert.NotEmpty(t, latest.Tags)
}

func TestUpdateExpenseByID(t *testing.T) {
	ex := seedExpense(t)

	body := bytes.NewBufferString(`{
		"title": "apple smoothie",
		"amount": 89,
		"note": "no discount", 
		"tags": ["beverage"]
	}`)
	var r Expense

	res := request(http.MethodPut, uri("expenses", strconv.Itoa(ex.ID)), body)
	err := res.Decode(&r)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.Equal(t, ex.ID, r.ID)
	assert.Equal(t, "apple smoothie", r.Title)
	assert.Equal(t, 89.0, r.Amount)
	assert.Equal(t, "no discount", r.Note)
	assert.Equal(t, []string{"beverage"}, r.Tags)
	assert.NotEmpty(t, r.Title)
	assert.NotEmpty(t, r.Amount)
	assert.NotEmpty(t, r.Note)
	assert.NotEmpty(t, r.Tags)
}

func seedExpense(t *testing.T) *Expense {
	var ex Expense
	body := bytes.NewBufferString(`{
		"title": "strawberry smoothie",
		"amount": 79,
		"note": "night market promotion discount 10 bath",
		"tags": ["food", "beverage"]
	}`)
	err := request(http.MethodPost, uri("expenses"), body).Decode(&ex)
	if err != nil {
		t.Fatal("can't create expense:", err)
	}
	return &ex
}

func request(method, url string, body io.Reader) *Response {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "November 10, 2009")
	client := http.Client{}
	res, err := client.Do(req)
	return &Response{res, err}
}

type Response struct {
	*http.Response
	err error
}

func (r *Response) Decode(v interface{}) error {
	if r.err != nil {
		return r.err
	}
	return json.NewDecoder(r.Body).Decode(v)
}

func uri(paths ...string) string {
	host := "http://localhost:2565"
	if paths == nil {
		return host
	}

	url := append([]string{host}, paths...)
	return strings.Join(url, "/")
}
