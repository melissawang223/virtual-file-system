package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

// createFileCmd represents the createFileCmd command
var createFileCmd = &cobra.Command{
	Use:   "create-folder",
	Short: "This command will create a new folder",
	Long:  "This command will create a new folder",
	Run: func(cmd *cobra.Command, args []string) {
		var userName = "melissa_folder"

		if len(args) >= 1 && args[0] != "" {
			userName = args[0]
		}

		// create user
		fmt.Println("Success create folder:", userName)
	},
}

func init() {
	rootCmd.AddCommand(createFileCmd)
}
