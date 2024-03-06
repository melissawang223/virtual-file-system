package controller

import (
	"fmt"
	"virtualFileSystem/helper"
	"virtualFileSystem/model"
)

// list-folders [username] [--sort-name|--sort-created] [asc|desc]

func ListFileController(args []string) error {
	userName := ""
	folderName := ""
	sortType := ""
	sortDir := ""

	if len(args) >= 3 {
		userName = args[0]
		folderName = args[1]
		sortType = args[2]
		sortDir = args[3]
	} else {
		fmt.Println("Error: The Input is insufficient.")
		return
	}

	// check userName is valid or not
	if err := helper.CheckUserName(userName); err != nil {
		fmt.Println("Error: The User Input is not ok.")
		return
	}

	// check if this user exist
	if !model.UserExist(userName) {
		fmt.Printf("Error: The %s does not exist.\n", userName)
		return
	}

	// check folderName is valid or not
	if err := helper.CheckFolderName(folderName); err != nil {
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

	model.ListFolder(userName, sortType, sortDir)

}
