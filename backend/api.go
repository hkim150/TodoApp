package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "todoapp"
	password = "password"
	dbname   = "todoapp"
)

var db *sql.DB

func init() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func GetHomePageHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

type todoItem struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}

func GetTodosHandler(c echo.Context) error {
	rows, err := db.Query("SELECT * FROM todo;")
	if err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "could not fetch data")
	}

	defer rows.Close()

	var todoItems []*todoItem
	for rows.Next() {
		t := new(todoItem)
		if err = rows.Scan(&t.Id, &t.Content); err != nil {
			fmt.Println(err)
			return c.String(http.StatusInternalServerError, "could not fetch data")
		}

		todoItems = append(todoItems, t)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "could not fetch data")
	}

	if len(todoItems) == 0 {
		return c.String(http.StatusNoContent, "no data found")
	}

	return c.JSON(http.StatusOK, todoItems)
}

type postTodoParams struct {
	Content string `json:"content"`
}

func PostTodoHandler(c echo.Context) error {
	var params postTodoParams
	err := c.Bind(&params)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	execString := fmt.Sprintf("INSERT INTO todo (content) VALUES ('%s')", params.Content)
	if _, err = db.Exec(execString); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "failed to insert data")
	}

	return c.String(http.StatusCreated, "success")
}

type patchTodoParams struct {
	Content string `json:"content"`
}

func PatchTodoHandler(c echo.Context) error {
	id := c.Param("id")
	var params patchTodoParams
	err := c.Bind(&params)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	execString := fmt.Sprintf("UPDATE todo SET content = '%s' WHERE id=%s", params.Content, id)
	if _, err = db.Exec(execString); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "failed to update data")
	}

	return c.String(http.StatusOK, "success")
}

func DeleteTodoHandler(c echo.Context) error {
	id := c.Param("id")
	if _, err := db.Exec("DELETE FROM todo WHERE id=$1", id); err != nil {
		fmt.Println(err)
		return c.String(http.StatusInternalServerError, "failed to delete data")
	}

	return c.String(http.StatusOK, "success")
}
