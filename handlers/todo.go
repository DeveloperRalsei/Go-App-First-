package handlers

import (
	"encoding/json"
	"errors"
	"os"
	"todo/env"
	"todo/models"
)

func GetAllTodos() ([]models.Todo, error) {
	f, err := os.Open(env.TodosFile)
	if err != nil {
		return []models.Todo{}, errors.New("Something went wrong while reading todos.json :(")
	}
	defer f.Close()

	decoder := json.NewDecoder(f)

	data := []models.Todo{}
	for decoder.More() {
		decoder.Decode(&data)
	}

	return data, nil
}

func AddNewTodo(todo models.Todo) {
	// todos, _ := GetAllTodos()
}

func GetTodoFromId(id int64) models.Todo {
	todos, _ := GetAllTodos()
	var foundTodo models.Todo

	for _, t := range todos {
		if t.Id == id {
			foundTodo = t
			break
		}
	}

	return foundTodo
}
