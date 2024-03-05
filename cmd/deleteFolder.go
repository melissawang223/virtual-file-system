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

	model.DeleteFolder(userName, folderName)

	fmt.Printf("Delete %s successfully.\n", folderName)
}
