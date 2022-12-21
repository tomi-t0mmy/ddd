package main

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
