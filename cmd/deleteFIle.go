package cmd

import (
	"fmt"
	"virtualFileSystem/model"
)

// delete-folder [username] [foldername]

func DeleteFile(args []string) {
	userName := "melissa"
	folderName := "melissa_folder"

	if len(args) >= 2 && args[0] != "" && args[1] != "" {
		userName = args[0]
		folderName = args[1]
	} else {
		return
	}

	// check user exist
	if _, ok := model.UsersMap[userName]; !ok {
		fmt.Printf("Error: The %s doesn't exist.\n", userName)
		return
	}

	// check if this user's folder exist
	if !model.FolderExist(userName, folderName) {
		return
	}

	model.DeleteFile(userName, folderName, folderName)

	fmt.Printf("Delete %s successfully.\n", folderName)
}
