package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	var r ITodoRepository = &InMemoryTodoRepository{[]Todo{}}
	e.GET("/todo", getTodoHandler(r))
	e.POST("/todo", createTodoHandler(r))

	e.Logger.Fatal(e.Start(":1323"))
}

func getTodoHandler(r ITodoRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		todos := r.getAll()
		return c.String(http.StatusOK, TodoPresenter{}.TodosPresent(todos))
	}
}

func createTodoHandler(r ITodoRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		todo, _ := NewTodo("hoge", false)
		r.createTodo(todo)
		return c.String(http.StatusOK, TodoPresenter{}.TodoPresent(*todo))
	}
}
