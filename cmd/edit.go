/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"github.com/subsavage/taskmaster/tasks"
	"github.com/spf13/cobra"
	"strconv"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit [task ID] [new title]",
	Short: "Edit the title of a task",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("❌ Invalid ID. Please enter a number.")
			return
		}

		newTitle := args[1]
		tasks.EditTask(id, newTitle)
	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
