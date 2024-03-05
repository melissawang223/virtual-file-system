package cmd

import (
	"fmt"
	"virtualFileSystem/helper"
	"virtualFileSystem/model"
)

// list-folders [username] [--sort-name|--sort-created] [asc|desc]

func ListFile(args []string) {
	userName := ""
	folderName := ""
	sortType := "--sort-name"
	sortDir := "asc"

	if len(args) >= 2 && args[0] != "" && args[1] != "" {
		userName = args[0]
		folderName = args[1]
		if len(args) >= 3 && args[2] != "" {
			sortType = args[2]
		}
		if len(args) >= 4 && args[3] != "" {
			sortDir = args[3]
		}
	} else {
		return
	}

	// check userName is valid or not
	if err := helper.CheckUser(userName); err != nil {
		return
	}

	// check if this user exist
	if model.UserExist(userName) {
		return
	}

	// check folderName is valid or not
	if err := helper.CheckFolder(folderName); err != nil {
		return
	}

	// check if this user's folder exist
	if !model.FolderExist(userName, folderName) {
		return
	}

	// check folder exist
	currentUser := model.UsersMap[userName]
	if len(currentUser.Folders) == 0 {
		fmt.Printf("Warning: The %s doesn't have any folders.\n", userName)
		return
	}

}
