package main

import (
	"path/filepath"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/Tasky/cmd"
	"github.com/spf13/Tasky/db"
)

func main() {
	home, _ := homedir.Dir()                  // Gets the homepath of the USER
	dbPath := filepath.Join(home, "tasks.db") // 'dbPath' contains the location where the daatbase will be created
	err := db.Init(dbPath)                    // Creates the database at 'dbPath'
	if err != nil {
		panic(err)
	}
	cmd.Execute()
} // Executes the main function
