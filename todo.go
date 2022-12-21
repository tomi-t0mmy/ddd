package main

import "errors"

type Todo struct {
	TodoText string
	Is_done  bool
}

func NewTodo(t string, d bool) (*Todo, error) {
	if t == "" {
		return nil, errors.New("タスク名を入力してください")
	}

	todo := &Todo{
		TodoText: t,
		Is_done:  false,
	}

	return todo, nil
}
