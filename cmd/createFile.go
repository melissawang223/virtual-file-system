package cmd

import (
	"fmt"
	"time"
	"virtualFileSystem/helper"
	"virtualFileSystem/model"
)

// create-file [username] [foldername] [filename] [description]?
func CreateFile(args []string) {

	userName := ""
	folderName := ""
	fileName := ""
	description := ""

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

	// check userName is valid or not
	if err := helper.CheckUser(userName); err != nil {
		return
	}

	// check if this user exist
	if !model.UserExist(userName) {
		return
	}

	// check folderName is valid or not
	if err := helper.CheckFolder(folderName); err != nil {
		return
	}

	// check if this user's folder exist
	if !model.FolderExist(userName, folderName) {
		return
	}

	//create folder
	currentUser := model.UsersMap[userName]
	if _, ok := currentUser.Folders[folderName]; !ok {
		fmt.Printf("Error: The %s doesn't exist.\n", userName)
		return
	}

	currentFolder := currentUser.Folders[folderName]
	currentFolder.File[fileName] = &model.File{
		Name:        fileName,
		Description: description,
		CreatedAt:   time.Now().Unix(),
		FolderName:  folderName,
		UserName:    userName,
	}

	fmt.Printf("Create %s in %s / %s  successfully.\n", fileName, userName, folderName)
}
