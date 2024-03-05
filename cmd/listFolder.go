package cmd

import (
	"fmt"
	"virtualFileSystem/model"
)

// list-folders [username] [--sort-name|--sort-created] [asc|desc]

func ListFolder(args []string) {
	userName := "melissa"
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

	// check user exist
	if _, ok := model.UsersMap[userName]; !ok {
		_ = fmt.Errorf("Error: The %s doesn't exist.\n", userName)
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
