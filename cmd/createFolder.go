package cmd

import (
	"fmt"
	"regexp"
	"time"
	"virtualFileSystem/file"
	"virtualFileSystem/folder"
	"virtualFileSystem/user"
)

// create-folder [username] [foldername] [description]

func CreateFolder(args []string) {
	userName := "melissa"
	folderName := "melissa_folder"
	description := "melissa_folder_description"

	if len(args) >= 2 && args[0] != "" && args[1] != "" {
		userName = args[0]
		folderName = args[1]
		if len(args) >= 3 && args[0] != "" {
			description = args[2]
		}
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

	//create folder
	currentUser := user.UsersMap[userName]
	currentUser.Folders[folderName] = &folder.Folder{
		Name:        folderName,
		Description: description,
		CreatedAt:   time.Now().Unix(),
		File:        map[string]*file.File{},
	}

	fmt.Printf("Create %s successfully.\n", folderName)
}
