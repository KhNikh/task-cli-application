package main

import (
	"fmt"
	"path/filepath"

	"github.com/KhNikh/task/cmd"
	"github.com/KhNikh/task/db"
	"github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir()
	dbPath := filepath.Join(home, "tasks.db")
	err := db.Init(dbPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("DB initialized")
	cmd.RootCmd.Execute()

}
