package main

import "errors"

type TodoToggleUsecase interface {
	call(int) (*Todo, error)
}

type TodoToggleInteractor struct {
	r ITodoRepository
}

func (i TodoToggleInteractor) call(id int) (*Todo, error) {
	todo, err := i.r.toggleDone(id)
	if err != nil {
		return nil, errors.New("error: the todo doesn't exist")
	}

	return todo, nil
}
