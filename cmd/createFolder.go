package cmd

import (
	"fmt"
	"time"
	"virtualFileSystem/helper"
	"virtualFileSystem/model"
)

// create-folder [username] [foldername] [description]

func CreateFolder(args []string) {
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

	// check folder description is valid or not
	if err := helper.CheckFileDescription(description); err != nil {
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
