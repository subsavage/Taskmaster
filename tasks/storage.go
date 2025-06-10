package tasks

import (
	"encoding/json"
	"os"
)

const dataFile = "tasks.json"

func SaveTasks() error {
	data, err := json.MarshalIndent(taskList, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, data, 0644)
}

func LoadTasks() error {
	file, err := os.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			taskList = []Task{}
			return nil
		}
		return err
	}
	return json.Unmarshal(file, &taskList)
}
