package controller

import (
	"fmt"
	"virtualFileSystem/helper"
	"virtualFileSystem/model"
)

// list-folders [username] [--sort-name|--sort-created] [asc|desc]

func ListFolderController(args []string) {
	userName := ""
	sortType := "--sort-name"
	sortDir := "asc"

	if len(args) >= 1 && args[0] != "" {
		userName = args[0]
		if len(args) >= 2 && args[1] != "" {
			sortType = args[1]
		}
		if len(args) >= 3 && args[2] != "" {
			sortDir = args[2]
		}
	} else {
		return
	}

	// check userName is valid or not
	if err := helper.CheckUserName(userName); err != nil {
		return
	}

	// check if this user exist
	if !model.UserExist(userName) {
		fmt.Printf("Error: The %s does not exist.\n", userName)
		return
	}

	// check folder exist
	currentUser := model.UsersMap[userName]
	if len(currentUser.Folders) == 0 {
		fmt.Printf("Warning: The %s doesn't have any folders.\n", userName)
		return
	}

	//list folder
	model.ListFolder(userName, sortType, sortDir)
}
