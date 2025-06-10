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

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done [task ID]",
	Short: "Mark a task as done",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("❌ Invalid ID. Please enter a number.")
			return
		}
		tasks.MarkDone(id)
	},
}


func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
