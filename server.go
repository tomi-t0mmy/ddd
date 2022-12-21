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

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

//e.GET("/show", show)
func show(c echo.Context) error {
	// Get team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

// e.POST("/save", save)
func save(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}
