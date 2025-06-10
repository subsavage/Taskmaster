package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"github.com/fatih/color"
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

func showTasks(filter ...string) {
	for _, task := range tasks {
		match := true

		if len(filter) > 0 {
			switch filter[0] {
			case "done":
				match = task.Status
			case "pending":
				match = !task.Status
			default:
				match = true
			}
		}

		if match {
			var state string
			var titleColored string

			if task.Status {
				state = color.HiGreenString("✅")
				titleColored = color.New(color.FgHiBlack).Sprint(task.Title)
			} else {
				state = color.HiRedString("❌")
				titleColored = color.New(color.FgHiWhite).Sprint(task.Title)
			}

			fmt.Printf("[%d] %s - %s\n", task.ID, state, titleColored)
		}
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

func deleteTask(id int) {
	index := -1

	for i, task := range tasks {
		if task.ID == id{
			index = i
			break
		}
	}

	if index == -1 {
		fmt.Println("Task not found, use the list command to see the tasks")
	}

	tasks = append(tasks[:index], tasks[index+1:]...)

	for i := range tasks {
        tasks[i].ID = i + 1
    }


	fmt.Printf("Deleted task #%d\n", id)
}

func editTask(id int, newTitle string) {
	for i := range tasks {
		if tasks[i].ID == id{
			tasks[i].Title = newTitle
			fmt.Printf("Task #%d updated successfully", id)
			return
		}
	}
	fmt.Println("Task ID not found.")
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


	case "delete":
		if len(args) < 3 {
			fmt.Println("Please provide the task ID")
			return
		}

		id,err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}

		deleteTask(id)
		err = saveTasks()

		if err!= nil {
			fmt.Println("Failed to save task: ", err)
		}

	case "edit":
		if len(args) < 4 {
			fmt.Println("Usage: edit <task ID> <new title>")
			return
		}

		id, err := strconv.Atoi(args[2])
		
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}

		newTitle := args[3]
		editTask(id, newTitle)

		err = saveTasks()

		if err != nil {
			fmt.Println("Failed to save tasks:", err)
    	}


	case "list":
	if len(args) == 3 && (args[2] == "done" || args[2] == "pending") {
		showTasks(args[2])
	} else {
		showTasks()
	}

	
	case "help":
		fmt.Println("HelpBook\n----------------------\nTo add a task, run : go run main.go add <your task>\nTo show all the tasks, run : go run main.go add list\nTo update the status of the task, run : go run main.go done <task number>\nTo delete a task, run : go run main.go delete <task number>\nTo edit a task, run : go run main.go edit <task number> <your new message>\nTo show pending tasks, run : go run main.go list pending\nTo show completed tasks, run : go run main.go list done\n----------------------")

	default:
		fmt.Println("Command not recognised, run the following command for help\n ")
	}

}