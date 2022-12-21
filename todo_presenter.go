package main

import (
	"encoding/json"
	"fmt"
)

type TodoPresenter struct {
}

func (p TodoPresenter) TodoPresent(t Todo) string {
	jsonData, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
		return "error!"
	}
	return string(jsonData)

}

func (p TodoPresenter) TodosPresent(todos []Todo) string {
	t := map[string][]Todo{"todos": todos}
	jsonData, err := json.Marshal(t)
	if err != nil {
		fmt.Println(err)
		return "error!"
	}
	return string(jsonData)
}
