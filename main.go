package main

import (
	// "todo/models"
	"fmt"
	"todo/handlers"
)

func main() {
	handlers.ListTasks()
	for {
		value := handlers.WaitForResponse()
		fmt.Printf("Girilen Değer: %s\n", value)
	}
}
