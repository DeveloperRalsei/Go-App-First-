package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"todo/handlers"
	"todo/models"
	"todo/utils"
)

func main() {
	Init()
	for {
	start:
		handlers.ListTasks()
		value := handlers.WaitForResponse()
		reader := bufio.NewReader(os.Stdin)

		switch value {
		case "1":
			todos, _ := handlers.GetAllTodos()

			utils.Clear()
			for _, t := range todos {
				fmt.Printf("%d[%s] - %s \n", t.Id, map[bool]string{
					true:  "x",
					false: " ",
				}[t.IsComplated], t.Name)
			}
			time.Sleep(3 * time.Second)
		case "2":
			utils.Clear()
			fmt.Printf("Enter the id of To-Do: ")
			todoStr, _ := reader.ReadString('\n')
			todoStr = strings.TrimSpace(todoStr)

			todoId, err := strconv.ParseInt(todoStr, 10, 64)
			if err != nil {
				fmt.Println("Invalid ID format. Please enter a valid number.")
				time.Sleep(2000 * time.Millisecond)
				utils.Clear()
				goto start
			}

			todo, todoErr := handlers.GetTodoFromId(todoId)
			if todoErr != nil {
				fmt.Printf("To-Do Not Found: %s", todoStr)
				time.Sleep(1500 * time.Millisecond)
				utils.Clear()
				goto start
			}

			handlers.ViewTodo(todo)
		case "3":
			utils.Clear()
			todos, _ := handlers.GetAllTodos()

			if len(todos) == 0 {
				if err := os.MkdirAll("data", 0755); err != nil {
					fmt.Printf("Couldn't create dir: %v\n", err)
					break
				}
				file, err := os.Create("data/todos.json")
				if err != nil {
					fmt.Printf("Couldn't create file: %v\n", err)
					break
				}
				file.Close()
			}

			var todo models.Todo

			fmt.Print("Please enter the name of To-Do (0 for exit): ")
			todoName, _ := reader.ReadString('\n')
			todoName = strings.TrimSpace(todoName)
			if todoName == "0" {
				utils.Clear()
				break
			}
			todo.Name = todoName

			fmt.Print("\nPlease enter a description: ")
			todoDesc, _ := reader.ReadString('\n')
			todo.Description = strings.TrimSpace(todoDesc)

			for {
				fmt.Print("\nIs this To-Do completed? (0-1): ")
				var isCompleted string
				fmt.Scan(&isCompleted)

				if isCompleted == "1" {
					todo.IsComplated = true
					break
				} else if isCompleted == "0" {
					todo.IsComplated = false
					break
				} else {
					fmt.Println("Invalid input. Please enter 0 or 1.")
				}
			}

			if err := handlers.AddNewTodo(&todo); err != nil {
				fmt.Printf("Error adding new To-Do: %v\n", err)
			}

		case "0":
			utils.Exit()

		default:
			utils.Clear()
			fmt.Println("Wrong usage\nPlease select a number between 1-4")
			time.Sleep(1500 * time.Millisecond)
			utils.Clear()
			handlers.ListTasks()
		}
	}
}
