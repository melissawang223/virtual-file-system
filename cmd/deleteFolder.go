package cmd

import (
	"fmt"
	"virtualFileSystem/helper"
	"virtualFileSystem/model"
)

// delete-folder [username] [foldername]

func DeleteFolder(args []string) {
	userName := ""
	folderName := ""

	if len(args) >= 2 && args[0] != "" && args[1] != "" {
		userName = args[0]
		folderName = args[1]
	} else {
		fmt.Println("Error: The Input is insufficient.")
		return
	}

	// check userName is valid or not
	if err := helper.CheckUser(userName); err != nil {
		return
	}

	// check if this user exist
	if !model.UserExist(userName) {
		fmt.Printf("Error: The User %s doesn't exist.\n", userName)
		return
	}

	// check folderName is valid or not
	if err := helper.CheckFolder(folderName); err != nil {
		return
	}

	// check if this user exist
	if !model.FolderExist(userName, folderName) {
		fmt.Printf("Error: The Folder %s doesn't exist.\n", folderName)
		return
	}

	model.DeleteFolder(userName, folderName)

	fmt.Printf("Delete %s successfully.\n", folderName)
}
