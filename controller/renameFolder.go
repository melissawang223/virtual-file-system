package controller

import (
	"fmt"
	"virtualFileSystem/helper"
	"virtualFileSystem/model"
)

func RenameFolderController(args []string) error {
	userName := ""
	folderName := ""
	newfolderName := ""

	if len(args) >= 3 {
		userName = args[0]
		folderName = args[1]
		newfolderName = args[2]
	} else {
		return fmt.Errorf("Usage: rename-folder [username] [foldername] [new-folder-name]`")
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

	// check newfolderName is valid or not
	if err := helper.CheckFolderName(newfolderName); err != nil {
		return err
	}

	// check if this user's folder name exist
	if !model.FolderExist(userName, folderName) {
		return fmt.Errorf("Error: The folder %s doesn't exist.\n", folderName)
	}

	// check if this user's new folder name exist
	if model.FolderExist(userName, newfolderName) {
		return fmt.Errorf("Error: The new-folder-name %s has already existed.\n", newfolderName)
	}

	model.ReNameFolder(userName, folderName, newfolderName)

	fmt.Printf("Rename %s to %s successfully.\n", folderName, newfolderName)

	return nil
}
