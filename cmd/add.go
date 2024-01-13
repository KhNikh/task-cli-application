package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/KhNikh/task/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Task is a cli task manager",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong")
			os.Exit(1)
		}
		fmt.Printf("added \"%s\" to your task \n", task)
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
	RootCmd.PersistentFlags().StringP("author", "a", "Nipun", "nipun")

}
