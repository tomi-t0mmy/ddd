package main

import "errors"

type InMemoryTodoRepository struct {
	maxId int
	todos []Todo
}

func (r *InMemoryTodoRepository) getAll() []Todo {
	return r.todos
}

func (r *InMemoryTodoRepository) createTodo(t *Todo) {
	t.Id = r.maxId + 1
	r.maxId = r.maxId + 1
	r.todos = append(r.todos, *t)
}

func (r *InMemoryTodoRepository) searchTodo(id int) (*Todo, error) {
	for _, todo := range r.todos {
		if todo.Id == id {
			return &todo, nil
		}
	}
	return nil, errors.New("error : The Todo doesn't exist")
}
