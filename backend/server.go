package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	e.GET("/", GetHomePageHandler)
	e.GET("/todos", GetTodosHandler)
	e.POST("/todos", PostTodoHandler)
	e.PATCH("/todos/:id", PatchTodoHandler)
	e.DELETE("/todos/:id", DeleteTodoHandler)

	e.Logger.Fatal(e.Start(":1323"))
}
