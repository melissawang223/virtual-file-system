package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"virtualFileSystem/user"
)

// createUserCmd represents the createUserCmd command
var createUserCmd = &cobra.Command{
	Use:   "register",
	Short: "This command will create a new user",
	Long:  "This command will create a new user",
	Run: func(cmd *cobra.Command, args []string) {
		var userName = "melissa"

		if len(args) >= 1 && args[0] != "" {
			userName = args[0]
		}

		// create user

		user.Users = append(user.Users, user.User{
			Name:    userName,
			Folders: map[string]string{},
		})

		fmt.Println("Success create user:", userName)
	},
}

func init() {
	rootCmd.AddCommand(createUserCmd)
}
