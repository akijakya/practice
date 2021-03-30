package cmd

import (
	"os"
	"github.com/akijakya/gophercises/task/db"
	"github.com/spf13/cobra"
	"fmt"
)

var listCmd = &cobra.Command{
    Use: "list",
    Short: "Lists all of your tasks.",
    Run: func(cmd *cobra.Command, args []string) {
        tasks, err := db.AllTasks()
        if err != nil {
            fmt.Println("Something went wrong: ", err)
            os.Exit(1)
        }
        if len(tasks) == 0 {
            fmt.Println("You have no tasks to complete!")
            return
        }
        fmt.Println("You have the following tasks:")
        for i, task := range tasks {
            fmt.Printf("%d. %s\n", i+1, task.Value)
        }
    },
}

// all init functions run before the main function in Go
func init() {
    RootCmd.AddCommand(listCmd)
}