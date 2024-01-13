package main

import (
	"path/filepath"

	"github.com/KhNikh/task/cmd"
	"github.com/KhNikh/task/db"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

func completionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "completion",
		Short: "Generate the autocompletion script for the specified shell",
	}
}

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	err := db.Init(dbPath)
	if err != nil {
		panic(err)
	}
	// fmt.Println("DB initialized")
	cmd.RootCmd.Execute()

}

func init() {
	completion := completionCommand()

	// mark completion hidden
	completion.Hidden = true
	cmd.RootCmd.AddCommand(completion)
}
