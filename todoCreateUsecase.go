package main

type TodoCreateUsecase interface {
	call(string) (*Todo, error)
}

type TodoCreateInteractor struct {
	r ITodoRepository
}

func (i TodoCreateInteractor) call(t string) (*Todo, error) {
	todo, err := NewTodo(t, false)
	if err != nil {
		return nil, err
	}
	i.r.createTodo(todo)
	return todo, nil

}
