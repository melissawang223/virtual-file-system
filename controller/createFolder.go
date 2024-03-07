package controller

import (
	"fmt"
	"os"
	"time"
	"virtualFileSystem/helper"
	"virtualFileSystem/model"
)

func CreateFolderController(args []string) error {
	userName := ""
	folderName := ""
	description := ""

	if len(args) >= 2 {
		userName = args[0]
		folderName = args[1]
		if len(args) >= 3 {
			description = args[2]
		}
	} else {
		return fmt.Errorf("usage: create-folder [username] [foldername] [description]?`\n")
	}

	// check userName is valid or not
	if err := helper.CheckUserName(userName); err != nil {
		return err
	}

	// check if this user exist
	if !model.UserExist(userName) {
		return fmt.Errorf("Error: The %s doesn't exist.\n", userName)
	}

	// check folderName is valid or not
	if err := helper.CheckFolderName(folderName); err != nil {
		return err
	}

	// check if this user's folder exist
	if model.FolderExist(userName, folderName) {
		return fmt.Errorf("Error: The %s already exist.\n", folderName)
	}

	// check folder description is valid or not
	if err := helper.CheckFileDescription(description); err != nil {
		return fmt.Errorf("Error: The Description %s is not valid: the length should less than 30.\n", description)

	}

	//create folder
	model.CreateFolder(userName, folderName, &model.Folder{
		Name:        folderName,
		Description: description,
		CreatedAt:   time.Now().Unix(),
		File:        map[string]*model.File{},
	})

	fmt.Fprintf(os.Stdout, "Create %s successfully.\n", folderName)
	return nil
}
