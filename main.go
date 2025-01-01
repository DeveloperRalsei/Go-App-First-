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
			fmt.Printf("\n")
		case "2":
			utils.Clear()
			fmt.Printf("Enter the id of To-Do: ")
			todoStr, _ := reader.ReadString('\n')

			todoId, err := strconv.ParseInt(todoStr, 10, 64)
			if err != nil {
				fmt.Println("Invalid ID format. Please enter a valid number.")
				time.Sleep(2000 * time.Millisecond)
				utils.Clear()
				goto start
			}
			todo := handlers.GetTodoFromId(todoId)
			fmt.Printf("todo: %v\n", todo)
		case "3":
			utils.Clear()
			var todo models.Todo

			fmt.Print("Please enter the name of To-Do(0 for exit): ")
			todoName, _ := reader.ReadString('\n')
			todoName = strings.TrimSpace(todoName)
			if todoName == "0" {
				utils.Clear()
				goto start
			}
			todo.Name = strings.TrimSpace(todoName)

			fmt.Print("\nPlease enter a description: ")
			todoDesc, _ := reader.ReadString('\n')
			todo.Description = strings.TrimSpace(todoDesc)

		checkIsComplated:
			fmt.Print("\nIs this To-Do complated?(0-1) ")
			var isComplated string
			fmt.Scan(&isComplated)
			if isComplated == "1" {
				todo.IsComplated = true
			} else if isComplated == "0" {
				todo.IsComplated = false
			} else {
				fmt.Printf("Wrong usage")
				time.Sleep(1500 * time.Millisecond)
				goto checkIsComplated
			}

			handlers.AddNewTodo(todo)
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
