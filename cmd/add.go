package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/Tasky/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ") //Concatenates the space separated words and stores it in 'task'
		_, err := db.CreateTask(task)   // Calls the CreateTask function which adds the task to a database
		if err != nil {
			fmt.Println("Something went wrong:", err) // Prints out the error if something goes wrong
			return
		}
		fmt.Printf("Added \"%s\" to your task list.\n", task) // Prints out if task is added succesfully
	},
}

func init() {
	/*This code adds the 'add' command using rootCmd */
	rootCmd.AddCommand(addCmd)
}
