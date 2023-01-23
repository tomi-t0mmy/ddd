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

func (r *InMemoryTodoRepository) toggleDone(id int) (*Todo, error) {
	for i, todo := range r.todos {
		if todo.Id == id {
			todo.IsDone = !todo.IsDone
			r.todos[i] = todo
			return &todo, nil
		}
	}
	return nil, errors.New("error : The Todo doesn't exist")
}

func (r *InMemoryTodoRepository) deleteTodo(id int) (*Todo, error) {

	for i, todo := range r.todos {
		if todo.Id == id {
			r.todos = r.todos[:i+copy(r.todos[i:], r.todos[i+1:])]
			return &todo, nil
		}
	}
	return nil, errors.New("error : The Todo doesn't exist")
}
