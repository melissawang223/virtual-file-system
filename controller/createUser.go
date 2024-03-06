package controller

import (
	"fmt"
	"os"
	"virtualFileSystem/helper"
	"virtualFileSystem/model"
)

// CreateUserController checks the input and create a user if everything looks good
func CreateUserController(args []string) error {
	var userName = ""

	if len(args) >= 1 {
		userName = args[0]
	} else {
		return fmt.Errorf(" Usage: register [username]")
	}

	// check userName is valid or not
	if err := helper.CheckUserName(userName); err != nil {
		return err
	}

	// check if this user exist
	if model.UserExist(userName) {
		return fmt.Errorf("Error: The %s has already existed.\n", userName)
	}

	model.CreateUser(userName)

	fmt.Fprintf(os.Stdout, "Add %s successfully.\n", userName)
	return nil
}
