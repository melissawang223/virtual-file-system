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
		return fmt.Errorf("Usage: `list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]`")
	}

	// check userName is valid or not
	if err := helper.CheckUserName(userName); err != nil {
		return err
	}

	// check if this user exist
	if !model.UserExist(userName) {
		return fmt.Errorf("Error: The User %s doesn't exist.\n", userName)
	}

	// check folderName is valid or not
	if err := helper.CheckFolderName(folderName); err != nil {
		return err
	}

	// check if this user's folder exist
	if !model.FolderExist(userName, folderName) {
		return
	}

	// check folder exist
	currentUser := model.UsersMap[userName]
	if len(currentUser.Folders) == 0 {
		return fmt.Errorf("Error: The folder %s doesn't exist.\n", folderName)
	}

	model.ListFolder(userName, sortType, sortDir)
	return nil
}
