package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/Tasky/db"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists out your current tasks",

	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks() //Stores the list of all tasks in 'tasks'

		if err != nil {
			fmt.Println("Something went wrong")
			os.Exit(1)
		}
		if len(tasks) == 0 {
			//Handles edge case of length(tasks) = 0
			fmt.Println("You have no tasks to complete!")
			return
		}
		fmt.Println("You have the following list of tasks.")
		for i, task := range tasks {
			//Iterates over occurences of 'tasks' and prints them
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
	},
}

func init() {
	/*This code adds the 'list' command using rootCmd */
	rootCmd.AddCommand(listCmd)
}
