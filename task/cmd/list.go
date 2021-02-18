package cmd

import "github.com/spf13/cobra"
import "fmt"

var listCmd = &cobra.Command{
    Use: "list",
    Short: "Lists all of your tasks.",
    Run: func(cmd *cobra.Command, args []string) {
        fmt.Println("list called")
    },
}

// all init functions run before the main function in Go
func init() {
    RootCmd.AddCommand(listCmd)
}