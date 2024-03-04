package cmd

import (
	"fmt"
	"regexp"
	"virtualFileSystem/folder"
	"virtualFileSystem/user"
)

// rename-folder [username] [foldername] [new-folder-name]

func RenameFolder(args []string) {
	userName := "melissa"
	folderName := "melissa_folder"
	newfolderName := "melissa_folder_description"

	if len(args) >= 3 && args[0] != "" && args[1] != "" && args[0] != "" {
		userName = args[0]
		folderName = args[1]
		newfolderName = args[2]
	} else {
		return
	}

	// check user exist
	if _, ok := user.UsersMap[userName]; !ok {
		fmt.Printf("Error: The %s doesn't exist.\n", userName)
		return
	}

	// check invalid char
	usernameConvention := "^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$"
	re, _ := regexp.Compile(usernameConvention)
	if len(folderName) > 40 || len(folderName) <= 0 || !re.MatchString(folderName) {
		fmt.Printf("Error: The %s contain invalid chars.", folderName)
		return
	}

	//rename folder
	currentUser := user.UsersMap[userName]
	oldFolder := currentUser.Folders[folderName]
	newFolder := &folder.Folder{
		Name:        newfolderName,
		Description: oldFolder.Description,
		CreatedAt:   oldFolder.CreatedAt,
		File:        oldFolder.File,
	}
	currentUser.Folders[newfolderName] = newFolder

	delete(currentUser.Folders, folderName)

	fmt.Printf("Rename %s to %s successfully.\n", folderName, newFolder)
}
