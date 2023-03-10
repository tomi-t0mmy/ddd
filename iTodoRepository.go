package main

type ITodoRepository interface {
	getAll() []Todo
	createTodo(*Todo)
	toggleDone(int) (*Todo, error)
	deleteTodo(int) (*Todo, error)
}
