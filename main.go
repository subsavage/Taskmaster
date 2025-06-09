package main

import (
	"fmt"
	"os"
)


type Task struct {
	ID int
	Title string
	Status bool
}


var tasks []Task

func addTask(title string){
	id := len(tasks)+1
	task := Task{ID:id, Title:title, Status:false}
	tasks = append(tasks, task)
}

func showTasks() {
	for _,task := range tasks {
		state := "❌"
		if task.Status{
			state = "✅"
		}
		fmt.Printf("[%d] %s - %s\n", task.ID, state, task.Title)
	}
}


func main() {
	args := os.Args

	if len(args)<2 {
		fmt.Println("Wrong Command, instead use\n go run main.go add <task>")
		return
	}

	switch args[1] {
	case "add":
		if len(args)<3 {
			fmt.Println("Looks like you forgot to mention the task title")
			return
		}
		title := args[2]
		addTask(title)
		fmt.Println("Task added to the list")
		showTasks()


	case "list":
		showTasks()
	
	case "help":
		fmt.Println("HelpBook\n To add a task, run : go run main.go add <your task>\n")

	default:
		fmt.Println("Command not recognised, run the following command for help\n ")
	}

}