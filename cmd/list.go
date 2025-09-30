/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/rival231/godo/tasks"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Use this to list down all the tasks",
	Long:  `Use this command to list down all your tasks in a table and see which are complete and incomplete`,
	Run: func(cmd *cobra.Command, args []string) {
		ts, err := tasks.LoadTasks()
		if err != nil {
			fmt.Println("Error loading tasks:", err)
			return
		}
		if len(ts) == 0 {
			fmt.Println("No tasks found.")
			return
		}

		// Print table header
		fmt.Printf("%-4s %-3s %-20s %-10s %-40s\n", "No.", "✓", "Task Name", "Priority", "Description")
		fmt.Println(strings.Repeat("-", 80))

		// Print tasks in rows
		for i, t := range ts {
			status := " "
			if t.Completed {
				status = "x"
			}
			taskName := t.Taskname
			if len(taskName) > 20 {
				taskName = taskName[:17] + "..." // truncate long task names
			}
			desc := t.Description
			if len(desc) > 40 {
				desc = desc[:37] + "..." // truncate long descriptions
			}
			fmt.Printf("%-4d [%-1s] %-20s %-10s %-40s\n", i+1, status, taskName, t.Priority, desc)
		}

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
