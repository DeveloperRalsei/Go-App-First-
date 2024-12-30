package main

import (
	"fmt"
	"time"
	"todo/handlers"
	"todo/models"
	"todo/utils"
)

func main() {
	Init()
	handlers.ListTasks()
	for {
		value := handlers.WaitForResponse()

		switch value {
		case "1":
			handlers.ListTasks()
			todos := handlers.GetAllTodos()

			utils.Clear()
			for _, todo := range todos {
				if todo.IsComplated {
					fmt.Printf("%d - %s | Complated\n", todo.Id, todo.Name)
				} else {
					fmt.Printf("%d - %s | Not Complated\n", todo.Id, todo.Name)
				}
			}
		case "2":
			handlers.ListTasks()
			var todo models.Todo

			handlers.AddNewTodo(todo)
		default:
			fmt.Println("Wrong usage")
			time.Sleep(1000)
		}
	}
}
