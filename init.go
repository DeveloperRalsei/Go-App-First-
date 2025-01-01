package main

import (
	"fmt"
	"todo/handlers"
	"todo/models"
	"todo/utils"
)

func Init() {
	handlers.Welcome()

	taskNames := []string{
		"List all To-Do Items",
		"View a To-Do Item",
		"Create a new To-Do item",
		"Delete a To-Do item",
		"Update a To-Do item",
	}

	for i, name := range taskNames {
		handlers.CreateNewTask(models.Task{
			Priority: i + 1,
			Name:     name,
		})
	}

}

func listTodosTask() {

	todos, _ := handlers.GetAllTodos()

	utils.Clear()
	for _, todo := range todos {
		fmt.Printf("%d - %s | %s ", todo.Id, todo.Name, map[bool]string{
			true:  "Complated\n",
			false: "Not Complated\n",
		}[todo.IsComplated])
	}
}
