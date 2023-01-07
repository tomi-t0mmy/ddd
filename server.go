package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	var inMemoryTodoRepository ITodoRepository = &InMemoryTodoRepository{0, []Todo{}}
	var todoCreateInteractor TodoCreateUsecase = &TodoCreateInteractor{inMemoryTodoRepository}
	var todoToggleInteractor TodoToggleUsecase = &TodoToggleInteractor{inMemoryTodoRepository}
	var todoDeleteInteractor TodoDeleteUsecase = &TodoDeleteInteractor{inMemoryTodoRepository}
	e.Use(middleware.CORS())
	e.GET("/todo", getTodoHandler(inMemoryTodoRepository))
	e.GET("/todo/:id", toggleDoneHandler(todoToggleInteractor))
	e.POST("/todo", createTodoHandler(todoCreateInteractor))
	e.DELETE("/todo/:id", deleteTodoHandler(todoDeleteInteractor))

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

func createTodoHandler(i TodoCreateUsecase) echo.HandlerFunc {
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

func toggleDoneHandler(i TodoToggleUsecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(strings.TrimRight(c.Param("id"), "\n"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid Parameters")
		}

		todo, err := i.call(id)
		if err != nil {
			return c.String(http.StatusBadRequest, "The Todo doesn't exist")
		}

		return c.String(http.StatusOK, TodoPresenter{}.TodoPresent(*todo))
	}
}

func deleteTodoHandler(i TodoDeleteUsecase) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(strings.TrimRight(c.Param("id"), "\n"))
		if err != nil {
			return c.String(http.StatusBadRequest, "Invalid Parameters")
		}

		todo, err := i.call(id)
		if err != nil {
			return c.String(http.StatusBadRequest, "The Todo doesn't exist")
		}

		return c.String(http.StatusOK, TodoPresenter{}.TodoPresent(*todo))
	}
}
