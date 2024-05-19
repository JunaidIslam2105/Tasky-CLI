package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/Tasky/db"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Marks a given task as completed and deletes it from the list",
	Run: func(cmd *cobra.Command, args []string) {
		var ids []int // Stores the ids of the tasks to be marked as done in a int slice 'ids'
		for _, arg := range args {
			//Iterates through each id
			id, err := strconv.Atoi(arg) // Attempts to parse the id and convert to int
			if err != nil {
				fmt.Println("Failed to parse the argument:", arg) //Prints error message if the id fails to be parseed
			} else {
				ids = append(ids, id) // Valid ids are appended to 'ids'
			}
		}
		tasks, err := db.AllTasks() // Gives us all the tasks in the database from ALLTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			return
		}
		for _, id := range ids {
			//Iterates over all valid integer ids
			if id <= 0 || id > len(tasks) {
				//checks if it is a valid id
				fmt.Println("Invalid task number:", id)
				continue
			}
			err := db.DeleteTask(tasks[id-1].Key) // Deletes the task corresponding to the task id
			if err != nil {
				fmt.Printf("Failed to mark \"%d\" as completed. Error: %s\n", id, err) // Prints error message
			} else {
				fmt.Printf("Marked \"%d\" as completed.\n", id) //Prints on succesful deletion
			}
		}
	},
}

func init() {
	/*This code adds the 'do' command using rootCmd */
	rootCmd.AddCommand(doCmd)
}
