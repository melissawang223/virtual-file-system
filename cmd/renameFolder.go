package cmd

import (
	"fmt"
	"virtualFileSystem/helper"
	"virtualFileSystem/model"
)

// rename-folder [username] [foldername] [new-folder-name]

func RenameFolderController(args []string) {
	userName := ""
	folderName := ""
	newfolderName := ""

	if len(args) >= 3 && args[0] != "" && args[1] != "" && args[2] != "" {
		userName = args[0]
		folderName = args[1]
		newfolderName = args[2]
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
		return
	}

	// check folderName is valid or not
	if err := helper.CheckFolder(folderName); err != nil {
		return
	}

	// check newfolderName is valid or not
	if err := helper.CheckFolder(newfolderName); err != nil {
		return
	}

	// check if this user's folder name exist
	if !model.FolderExist(userName, folderName) {
		return
	}

	// check if this user's new folder name exist
	if !model.FolderExist(userName, newfolderName) {
		return
	}

	model.ReNameFolder(userName, folderName, newfolderName)

	fmt.Printf("Rename %s to %s successfully.\n", folderName, newfolderName)
}
