/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "taskmaster",
	Short: "A simple CLI task manager",
	Long: `HelpBook
----------------------
To add a task, run:            taskmaster add <your task>
To show all tasks, run:        taskmaster list
To update a task's status:     taskmaster done <task number>
To delete a task, run:         taskmaster delete <task number>
To edit a task, run:           taskmaster edit <task number> <your new message>
To show pending tasks, run:    taskmaster list pending
To show completed tasks, run:  taskmaster list done
----------------------`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.github.com/subsavage/taskmaster.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


