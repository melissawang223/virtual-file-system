package controller

import (
	"fmt"
	"time"
	"virtualFileSystem/helper"
	"virtualFileSystem/model"
)

func CreateFileController(args []string) error {

	userName := ""
	folderName := ""
	fileName := ""
	description := ""

	if len(args) >= 3 {
		userName = args[0]
		folderName = args[1]
		fileName = args[2]
		if len(args) >= 4 {
			description = args[3]
		}
	} else {
		return fmt.Errorf("Usage: create-file [username] [foldername] [filename] [description]?")
	}

	// check userName is valid or not
	if err := helper.CheckUserName(userName); err != nil {
		return err
	}

	// check if this user exist
	if !model.UserExist(userName) {
		return fmt.Errorf("Error: The User %s doesn't exist.\n", userName)
	}

	// check folderName is valid or not
	if err := helper.CheckFolderName(folderName); err != nil {
		return err
	}

	// check if this user's folder exist
	if !model.FolderExist(userName, folderName) {
		return fmt.Errorf("Error: The Folder %s does not exist.\n", folderName)
	}

	// check folderName is valid or not
	if err := helper.CheckFileName(fileName); err != nil {
		return err
	}

	// check if this user's folder exist
	if model.FileExist(userName, folderName, fileName) {
		return fmt.Errorf("Error: The File %s has already existed.\n", fileName)
	}

	//create folder
	newFile := &model.File{
		Name:        fileName,
		Description: description,
		CreatedAt:   time.Now().Unix(),
		FolderName:  folderName,
		UserName:    userName,
	}
	model.CreateFile(userName, folderName, fileName, newFile)

	fmt.Printf("Create %s in %s / %s successfully.\n", fileName, userName, folderName)
	return nil
}
