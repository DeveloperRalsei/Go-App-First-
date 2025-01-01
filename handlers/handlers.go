package handlers

import (
	"fmt"
	"os"
	"todo/models"
)

func WaitForResponse() string {
	var value string
	fmt.Printf("Please enter a value between 1 and 4 (0 for exit): ")
	fmt.Scan(&value)

	return value
}

func Welcome() {
	fmt.Println("--- Welcome to To-Do App ---")
}

func ListTasks() {
	fmt.Printf("\n")

	for i, task := range models.Tasks {
		fmt.Printf("%d - %v \n", i+1, task.Name)
	}
}

func CreateNewTask(task models.Task) []models.Task {
	models.Tasks = append(models.Tasks, task)
	return models.Tasks
}

func Exit() {
	os.Exit(1)
}
