package cmd

import (
	"fmt"
	"virtualFileSystem/user"
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
	if _, ok := user.UsersMap[userName]; !ok {
		fmt.Printf("Error: The %s doesn't exist.\n", userName)
		return
	}

	// check folder exist
	currentUser := user.UsersMap[userName]
	if _, ok := currentUser.Folders[folderName]; !ok {
		fmt.Printf("Error: The %s doesn't exist.\n", folderName)
		return
	}

	//delete folder
	delete(currentUser.Folders, folderName)

	fmt.Printf("Delete %s successfully.\n", folderName)
}