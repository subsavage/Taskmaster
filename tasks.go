package main

import (
	"fmt"
	"github.com/fatih/color"
)

type Task struct {
	ID     int
	Title  string
	Status bool
}

var tasks []Task

func addTask(title string) {
	id := len(tasks) + 1
	task := Task{ID: id, Title: title, Status: false}
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

func markDone(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = true
			fmt.Printf("Status of Task %d updated successfully\n", id)
			return
		}
	}
	fmt.Println("Task ID not found.")
}

func deleteTask(id int) {
	index := -1
	for i, task := range tasks {
		if task.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		fmt.Println("Task not found.")
		return
	}
	tasks = append(tasks[:index], tasks[index+1:]...)
	for i := range tasks {
		tasks[i].ID = i + 1
	}
	fmt.Printf("Deleted task #%d\n", id)
}

func editTask(id int, newTitle string) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Title = newTitle
			fmt.Printf("Task #%d updated successfully\n", id)
			return
		}
	}
	fmt.Println("Task ID not found.")
}
