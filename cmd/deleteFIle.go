package cmd

import (
	"fmt"
	"virtualFileSystem/helper"
	"virtualFileSystem/model"
)

func DeleteFileController(args []string) {
	userName := ""
	folderName := ""
	fileName := ""

	if len(args) >= 3 && args[0] != "" && args[1] != "" && args[2] != "" {
		userName = args[0]
		folderName = args[1]
		fileName = args[2]
	} else {
		fmt.Println("Error: The Input is insufficient.")
		return
	}

	// check userName is valid or not
	if err := helper.CheckUser(userName); err != nil {
		return
	}

	// check if this user exist
	if !model.UserExist(userName) {
		fmt.Printf("Error: The User %s doesn't exist.\n", userName)
		return
	}

	// check folderName is valid or not
	if err := helper.CheckFolder(folderName); err != nil {
		return
	}

	// check if this user exist
	if !model.FolderExist(userName, folderName) {
		fmt.Printf("Error: The Folder %s doesn't exist.\n", folderName)
		return
	}

	// check fileName is valid or not
	if err := helper.CheckFile(fileName); err != nil {
		return
	}

	// check if this user exist
	if !model.FileExist(userName, folderName, fileName) {
		fmt.Printf("Error: The File %s doesn't exist.\n", fileName)
		return
	}

	model.DeleteFile(userName, folderName, folderName)

	fmt.Printf("Delete %s in %s / %s successfully.\n", fileName, userName, folderName)
}
