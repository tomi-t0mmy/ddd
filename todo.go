package main

import "errors"

type Todo struct {
	Id       int
	TodoText string
	IsDone   bool
}

func NewTodo(t string, d bool) (*Todo, error) {
	if t == "" {
		return nil, errors.New("タスク名を入力してください")
	}

	todo := &Todo{
		Id:       -1, // repositoryで設定するため、一旦-1を入れておく
		TodoText: t,
		IsDone:   false,
	}

	return todo, nil
}
