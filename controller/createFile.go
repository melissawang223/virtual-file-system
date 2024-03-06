package controller

import (
	"fmt"
	"time"
	"virtualFileSystem/helper"
	"virtualFileSystem/model"
)

func CreateFileController(args []string) {

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
		fmt.Println("Error: The Input is insufficient.")
		return
	}

	// check userName is valid or not
	if err := helper.CheckUserName(userName); err != nil {
		return
	}

	// check if this user exist
	if !model.UserExist(userName) {
		fmt.Printf("Error: The User %s does not exist.\n", userName)
		return
	}

	// check folderName is valid or not
	if err := helper.CheckFolderName(folderName); err != nil {
		return
	}

	// check if this user's folder exist
	if !model.FolderExist(userName, folderName) {
		fmt.Printf("Error: The Folder %s does not exist.\n", folderName)
		return
	}

	// check folderName is valid or not
	if err := helper.CheckFileName(fileName); err != nil {
		return
	}

	// check if this user's folder exist
	if model.FileExist(userName, folderName, fileName) {
		fmt.Printf("Error: The Folder %s has already existed.\n", fileName)
		return
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
}
