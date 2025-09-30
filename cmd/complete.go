/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/rival231/godo/tasks"
	"github.com/spf13/cobra"
)

// completeCmd represents the complete command
var completeCmd = &cobra.Command{
	Use:   "complete",
	Short: "This is to mark task as completed",
	Long: `Mark the task as completed by providing the name of the task you want to mark as completed`,
	Run: func(cmd *cobra.Command, args []string) {
		if taskname == "" {
			fmt.Println("Task name cannot be empty")
		}
		ts, err := tasks.LoadTasks()
		if err != nil {
			fmt.Println("unable to load tasks")
		}
		found := false
		for index, task := range ts {
			if taskname == task.Taskname {
				ts[index].Completed = true
				err = tasks.SaveTasks(ts)
				if err != nil {
					fmt.Println("Failed to mark task as completed")
					return
				}
				found = true
				break
			}
		}

		if found {
			fmt.Println("Task marked completed")
		} else {
			fmt.Println("Task not found")
		}

	},
}

func init() {
	rootCmd.AddCommand(completeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")
	completeCmd.Flags().StringVarP(&taskname, "name", "t", "", "Enter the name of the to mark as complete")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
