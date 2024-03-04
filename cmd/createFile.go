package cmd

import (
	"fmt"
	"regexp"
	"time"
	"virtualFileSystem/file"
	"virtualFileSystem/user"
)

// create-file [username] [foldername] [filename] [description]?
func CreateFile(args []string) {

	userName := "melissa"
	folderName := "melissa_folder"
	fileName := "melissa_folder_file"
	description := "melissa_folder_description"

	if len(args) >= 3 && args[0] != "" && args[1] != "" && args[2] != "" {
		userName = args[0]
		folderName = args[1]
		fileName = args[2]
		if len(args) >= 4 && args[3] != "" {
			description = args[3]
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
	if _, ok := currentUser.Folders[folderName]; !ok {
		fmt.Printf("Error: The %s doesn't exist.\n", userName)
		return
	}

	currentFolder := currentUser.Folders[folderName]
	currentFolder.File[fileName] = &file.File{
		Name:        fileName,
		Description: description,
		CreatedAt:   time.Now().Unix(),
		FolderName:  folderName,
		UserName:    userName,
	}

	fmt.Printf("Create %s in %s / %s  successfully.\n", fileName, userName, folderName)
}
