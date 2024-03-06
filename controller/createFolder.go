package controller

import (
	"fmt"
	"time"
	"virtualFileSystem/helper"
	"virtualFileSystem/model"
)

func CreateFolderController(args []string) {
	userName := ""
	folderName := ""
	description := ""

	if len(args) >= 2 && args[0] != "" && args[1] != "" {
		userName = args[0]
		folderName = args[1]
		if len(args) >= 3 && args[2] != "" {
			description = args[2]
		}
	} else {
		_ = fmt.Errorf("Usage: create-folder [username] [foldername] [description]?`")
		return
	}

	// check userName is valid or not
	if err := helper.CheckUserName(userName); err != nil {
		return
	}

	// check if this user exist
	if !model.UserExist(userName) {
		_ = fmt.Errorf("Error: The %s doesn't exist.\n", userName)
		return
	}

	// check folderName is valid or not
	if err := helper.CheckFolderName(folderName); err != nil {
		return
	}

	// check if this user's folder exist
	if model.FolderExist(userName, folderName) {
		_ = fmt.Errorf("Error: The %s already exist.\n", folderName)
		return
	}

	// check folder description is valid or not
	if err := helper.CheckFileDescription(description); err != nil {
		_ = fmt.Errorf("Error: The Description %s is not valid: the length should less than 30.\n", description)
		return
	}

	//create folder
	model.CreateFolder(userName, folderName, &model.Folder{
		Name:        folderName,
		Description: description,
		CreatedAt:   time.Now().Unix(),
		File:        map[string]*model.File{},
	})

	fmt.Printf("Create %s successfully.\n", folderName)
}
