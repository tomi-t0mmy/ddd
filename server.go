package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	var inMemoryTodoRepository ITodoRepository = &InMemoryTodoRepository{0, []Todo{}}
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

type CreateTodoParam struct {
	Todo string `json:"todo"`
}

func createTodoHandler(r ITodoRepository, i TodoCreateUsecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		p := new(CreateTodoParam)

		if err := c.Bind(p); err != nil {
			fmt.Printf("err %v", err.Error())
			return c.String(http.StatusInternalServerError, "Error!")
		}

		todo, err := i.call(p.Todo)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid Parameters")
		}
		return c.String(http.StatusOK, TodoPresenter{}.TodoPresent(*todo))
	}
}
