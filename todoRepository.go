package main

type InMemoryTodoRepository struct {
	todos []Todo
}

func (r *InMemoryTodoRepository) getAll() []Todo {
	return r.todos
}

func (r *InMemoryTodoRepository) createTodo(t *Todo) {
	r.todos = append(r.todos, *t)
}
