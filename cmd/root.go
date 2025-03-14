package cmd

// Copyright © 2025

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"todo-cli/tasks"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "CLI to manage your tasks",
}

var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Add a new task",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_taskName := strings.Join(args, " ")
		tasks.AddTask(_taskName)
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	Run: func(cmd *cobra.Command, args []string) {
		tasks.DisplayList()
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete [taskId]",
	Short: "Delete a task",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_taskId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}

		tasks.DeleteTask(_taskId)
	},
}

var completeCmd = &cobra.Command{
	Use:   "complete [taskId]",
	Short: "Mark a task as done",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		_taskId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println(err)
			return
		}
		tasks.MarkDone(_taskId)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(addCmd, listCmd, deleteCmd, completeCmd)
}
