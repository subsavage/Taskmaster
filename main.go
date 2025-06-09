package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
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

func saveTasks() error{
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, data, 0644)
}

func loadTasks() error {
	file, err := os.ReadFile(dataFile)
	if err != nil{
		if os.IsNotExist(err) {
			tasks = []Task{}
			return nil
		}
		return err
	}
	return json.Unmarshal(file, &tasks)
}

func markDone(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = true
			fmt.Printf("Status of Task %d updated successfully", id)
			return
		}
	}
	fmt.Println("Task ID not found, use the list command to see all the tasks")
}

const dataFile = "tasks.json"

func main() {
	err := loadTasks()

	if err != nil{
		fmt.Println("Failed to load tasks: ", err)
		return
	}

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
		err := saveTasks()
		
		if err != nil {
			fmt.Println("Failed to save task: ", err)
		}

		fmt.Println("Task added to the list")
		showTasks()

	case "done":
		if len(args) < 3 {
			fmt.Println("Please provide the task ID")
			return
		}

		id, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}

		markDone(id)
		err = saveTasks()
		if err != nil {
			fmt.Println("Failed to save tasks: ", err)
		}


	case "list":
		showTasks()
	
	case "help":
		fmt.Println("HelpBook\n----------------------\nTo add a task, run : go run main.go add <your task>\nTo show all the tasks, run : go run main.go add list\n----------------------")

	default:
		fmt.Println("Command not recognised, run the following command for help\n ")
	}

}