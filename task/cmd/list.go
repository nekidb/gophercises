package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your incomplete tasks",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("You have not any tasks")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
