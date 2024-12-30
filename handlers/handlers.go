package handlers

import (
	"fmt"
	"todo/models"
)

func WaitForResponse() string {
	var value string
	fmt.Scan(&value)

	return value
}

func ListTasks() {
	fmt.Println("--- Welcome to To-Do App ---")
	fmt.Println("")

	for i, task := range models.Tasks {
		fmt.Printf("%d. %v", i+1, task.Name)
	}
}

func CreateNewTask(task models.Task) {
	models.Tasks = append(models.Tasks, task)

}
