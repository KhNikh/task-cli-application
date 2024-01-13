/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/KhNikh/task/db"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all your tasks",

	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong: ", err)
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no tasks to complete. Enjoy !!!")
			return
		}
		fmt.Println("You have the following tasks to complete")
		for i, task := range tasks {
			fmt.Printf("%d. %s \n", i+1, task.Value)
		}
	},
}

func init() {
	RootCmd.AddCommand(listCmd)

}
