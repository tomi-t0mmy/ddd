package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	var inMemoryTodoRepository ITodoRepository = &InMemoryTodoRepository{[]Todo{}}
	var todoCreateInteractor TodoCreateUsecase = &TodoCreateInteractor{inMemoryTodoRepository}
	e.GET("/todo", getTodoHandler(inMemoryTodoRepository))
	e.POST("/todo", createTodoHandler(inMemoryTodoRepository, todoCreateInteractor))

	e.Logger.Fatal(e.Start(":1323"))
}

func getTodoHandler(r ITodoRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		todos := r.getAll()
		return c.String(http.StatusOK, TodoPresenter{}.TodosPresent(todos))
	}
}

func createTodoHandler(r ITodoRepository, i TodoCreateUsecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		todo, err := i.call("hoge")
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid Parameters")
		}
		return c.String(http.StatusOK, TodoPresenter{}.TodoPresent(*todo))
	}
}
