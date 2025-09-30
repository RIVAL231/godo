/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import "github.com/rival231/godo/cmd"

type Task struct {
	Taskname    string 
	Description string
	priority    string
	Completed   bool
}

func main() {
	cmd.Execute()
}
