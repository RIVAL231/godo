/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/rival231/godo/tasks"
	"github.com/spf13/cobra"
)
var priority string
var taskname string
var taskdesc string

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "This command is to add tasks,",
	Long: `Use this tool to add tasks to your daily to do with priorities as high, low, or medium`,
	Run: func(cmd *cobra.Command, args []string) {
		
		ts,err := tasks.LoadTasks()
		if err!=nil{
			fmt.Println("error loading tasks")
		}
		t := tasks.Tasks{
			Taskname: taskname,
			Description: taskdesc,
			Priority: priority,
			Completed: false,
		}
		ts = append(ts, t)
		err = tasks.SaveTasks(ts)
        if err!=nil{
			fmt.Println("error saving task")
		}
		fmt.Println("task added successfully")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")
    addCmd.Flags().StringVarP(&priority,"priority","p","low","Set the priority for your task")
	addCmd.Flags().StringVarP(&taskname,"taskname","n","","Name of your task")
	addCmd.Flags().StringVarP(&taskdesc,"taskdesc","d","","Description of your task being added")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
