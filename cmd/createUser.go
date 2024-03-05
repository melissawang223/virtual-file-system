package cmd

import (
	"fmt"
	"virtualFileSystem/helper"
	"virtualFileSystem/model"
)

// CreateUserController checks the input and create a user if everything looks good
func CreateUserController(args []string) {
	var userName = ""

	if len(args) >= 1 && args[0] != "" {
		userName = args[0]
	} else {
		fmt.Println("Error: The Input is insufficient.")
		return
	}

	// check userName is valid or not
	if err := helper.CheckUser(userName); err != nil {
		return
	}

	// check if this user exist
	if model.UserExist(userName) {
		fmt.Printf("Error: The %s has already existed.\n", userName)
		return
	}

	model.CreateUser(userName)

	fmt.Printf("Add %s successfully.\n", userName)
}
