package controller

import (
	"fmt"
	"os"
	"virtualFileSystem/helper"
	"virtualFileSystem/model"
)

// delete-folder [username] [foldername]

func DeleteFolderController(args []string) error {
	userName := ""
	folderName := ""

	if len(args) >= 2 {
		userName = args[0]
		folderName = args[1]
	} else {
		return fmt.Errorf("Usage: delete-folder [username] [foldername]")
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

	// check if this user exist
	if !model.FolderExist(userName, folderName) {
		return fmt.Errorf("Error: The Folder %s doesn't exist.\n", folderName)
	}

	model.DeleteFolder(userName, folderName)

	fmt.Fprintf(os.Stdout, "Delete %s successfully.\n", folderName)
	return nil
}
