/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/subsavage/taskmaster/tasks"
	"github.com/spf13/cobra"
	"fmt"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list [all|done|pending]",
	Short: "List all tasks or filter by status",
	Run: func(cmd *cobra.Command, args []string) {
		err := tasks.LoadTasks() 
		if err != nil {
			fmt.Println("Failed to load tasks:", err)
			return
		}
		tasks.ShowTasks(args...) 
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
