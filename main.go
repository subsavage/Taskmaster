package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	err := loadTasks()
	if err != nil {
		fmt.Println("Failed to load tasks:", err)
		return
	}

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Wrong Command. Try:\n go run main.go add <task>")
		return
	}

	switch args[1] {
	case "add":
		if len(args) < 3 {
			fmt.Println("Missing task title")
			return
		}
		addTask(args[2])
		saveTasks()
		showTasks()

	case "done":
		if len(args) < 3 {
			fmt.Println("Provide the task ID")
			return
		}
		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}
		markDone(id)
		saveTasks()

	case "delete":
		if len(args) < 3 {
			fmt.Println("Provide the task ID")
			return
		}
		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}
		deleteTask(id)
		saveTasks()

	case "edit":
		if len(args) < 4 {
			fmt.Println("Usage: edit <id> <new title>")
			return
		}
		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}
		editTask(id, args[3])
		saveTasks()

	case "list":
		if len(args) == 3 && (args[2] == "done" || args[2] == "pending") {
			showTasks(args[2])
		} else {
			showTasks()
		}

	case "help":
		fmt.Println("HelpBook\n----------------------\nTo add a task, run : go run main.go add <your task>\nTo show all the tasks, run : go run main.go add list\nTo update the status of the task, run : go run main.go done <task number>\nTo delete a task, run : go run main.go delete <task number>\nTo edit a task, run : go run main.go edit <task number> <your new message>\nTo show pending tasks, run : go run main.go list pending\nTo show completed tasks, run : go run main.go list done\n----------------------")

	default:
		fmt.Println("Command not recognised. Use 'help' to see options.")
	}
}




