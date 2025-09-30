/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/rival231/godo/tasks"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a task",
	Long:  `Use this to delete a task by providing the name of the task`,
	Run: func(cmd *cobra.Command, args []string) {
		if taskname == "" {
			fmt.Println("Task name cannot be empty")
		}
		ts, err := tasks.LoadTasks()
		if err != nil {
			fmt.Println("Cannot load tasks")
		}
		var found bool
		for index, task := range ts {
			if taskname == task.Taskname {
				found = true
				ts = append(ts[:index], ts[index+1:]...)
				err = tasks.SaveTasks(ts)
				if err != nil {
					fmt.Println("Failed to delete task cannot save")
				}
			}

		}
		if found {
			fmt.Println("task deleted successfully")
		} else {
			fmt.Println("Task not found cannot delete")
		}
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")
	deleteCmd.Flags().StringVarP(&taskname, "taskname", "n", "", "Write name of task to delete")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
