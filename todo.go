package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    fmt.Println("Welcome to the TO DO List CLI app!")

	var command string
	var tasks []string
    
    for {
		fmt.Println()
		fmt.Println("Enter your command (create, read, update, delete):") 
		fmt.Scan(&command)

		switch command {
        case "create":
			fmt.Println("Enter task name:")
			task, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			tasks = append(tasks, task)
        case "read":
			for i, item := range tasks {
				fmt.Printf("%d. %s", i + 1, item)
			}
        case "update":
			fmt.Println("Enter task name to update:")
			task, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			
			for i, item := range tasks {
				if task == item {
					fmt.Println("Enter new task name:")
					tasks[i], _ = bufio.NewReader(os.Stdin).ReadString('\n')
					fmt.Printf("Update task #%d with name %s successfully", i, task)
				} else {
					fmt.Println("Invalid command! Please, try again!")
				}
			}
        
        case "delete":
            fmt.Println("Enter task name to remove:")
			task, _ := bufio.NewReader(os.Stdin).ReadString('\n')
			
			for i, item := range tasks {
				if task == item {
					tasks = append(tasks[:i], tasks[i+1:]...)
					fmt.Printf("Removed task #%d with name %s successfully\n", i, task)
				} else {
					fmt.Println("Invalid command! Please, try again!")
				}
			}
        default: 
            fmt.Println("Invalid command! Please, try again!")
		}
	}
}