package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
    Use: "do",
    Short: "Marks a task to complete.",
    Run: func(cmd *cobra.Command, args []string) {
        var ids []int
		for _, arg := range args {
			// go has a strconv.ParseUInt also, but it is designed to be more generic, 
			// to be able to parse int from any format
			// Atoi is ascii to integer, it does ParseInt in the background but with a few basic settings
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument: ", arg)
			} else {
				ids = append(ids, id)
			}
		}
		fmt.Println(ids)
    },
}

// all init functions run before the main function in Go
func init() {
    RootCmd.AddCommand(doCmd)
}