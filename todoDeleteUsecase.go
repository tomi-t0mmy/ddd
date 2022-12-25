package main

import "errors"

type TodoDeleteUsecase interface {
	call(int) (*Todo, error)
}

type TodoDeleteInteractor struct {
	r ITodoRepository
}

func (i TodoDeleteInteractor) call(id int) (*Todo, error) {
	todo, err := i.r.deleteTodo(id)
	if err != nil {
		return nil, errors.New("error: the todo doesn't exist")
	}

	return todo, nil
}
