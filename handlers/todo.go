package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
	"todo/env"
	"todo/models"
	"todo/utils"
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

func AddNewTodo(todo *models.Todo) error {
	todos, geterr := GetAllTodos()

	if geterr != nil {
		fmt.Printf("err: %v\n", geterr)
		return geterr
	}

	file, fileErr := os.OpenFile(env.TodosFile, os.O_RDWR|os.O_CREATE, 0644)
	defer file.Close()
	if fileErr != nil {
		return errors.New("Something went wrong while opening file :(")
	}

	todo.Id = int64(len(todos) + 1)
	newTodosList := append(todos, *todo)

	if encodeErr := json.NewEncoder(file).Encode(newTodosList); encodeErr != nil {
		return errors.New("Something went wrong while writing new Item :(")
	}
	return nil
}

func GetTodoFromId(id int64) (models.Todo, error) {
	todos, _ := GetAllTodos()
	var foundTodo models.Todo
	found := false

	for _, t := range todos {
		if t.Id == id {
			foundTodo = t
			found = true
			break
		}
	}

	if !found {
		return models.Todo{}, errors.New("Not Found")
	}

	return foundTodo, nil
}

func ViewTodo(todo models.Todo) {
	utils.Clear()
	fmt.Printf("To-Do Name: %s\n", todo.Name)
	fmt.Printf("To-Do Description: %s\n", todo.Description)
	fmt.Printf("Status: %s\n", map[bool]string{
		true:  "Complated",
		false: "Not Complated",
	}[todo.IsComplated])
	time.Sleep(1800 * time.Millisecond)
}
