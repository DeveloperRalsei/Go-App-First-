package main

import (
	"todo/handlers"
	"todo/models"
)

func Init() {
	handlers.Welcome()

	taskNames := []string{
		"List All To-Do's",
		"Create a new To-Do",
		"Update a To-Do",
		"Delete a To-do",
	}

	for _, taskName := range taskNames {
		handlers.CreateNewTask(models.Task{
			Name: taskName,
		})
	}
}
