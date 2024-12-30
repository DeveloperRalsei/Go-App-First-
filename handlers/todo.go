package handlers

import (
	"encoding/json"
	"os"
	"todo/env"
	"todo/models"
	"todo/utils"
)

func GetAllTodos() []models.Todo {
	f, _ := os.Open(env.TodosFile)
	defer f.Close()

	decoder := json.NewDecoder(f)

	data := []models.Todo{}
	for decoder.More() {
		decoder.Decode(&data)
	}

	return data
}

func AddNewTodo(todo models.Todo) {
	utils.Clear()
	return
}
