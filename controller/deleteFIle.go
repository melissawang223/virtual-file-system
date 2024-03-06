package controller

import (
	"fmt"
	"virtualFileSystem/helper"
	"virtualFileSystem/model"
)

func DeleteFileController(args []string) error {
	userName := ""
	folderName := ""
	fileName := ""

	if len(args) >= 3 {
		userName = args[0]
		folderName = args[1]
		fileName = args[2]
	} else {
		return fmt.Errorf("Usage: `delete-file [username] [foldername] [filename]`")
	}

	// check userName is valid or not
	if err := helper.CheckUserName(userName); err != nil {
		return err
	}

	// check if this user exist
	if !model.UserExist(userName) {
		return fmt.Errorf("Error: The %s has already existed.\n", userName)
	}

	// check folderName is valid or not
	if err := helper.CheckFolderName(folderName); err != nil {
		return err
	}

	// check if this user exist
	if !model.FolderExist(userName, folderName) {
		return fmt.Errorf("Error: The Folder %s doesn't exist.\n", folderName)
	}

	// check fileName is valid or not
	if err := helper.CheckFileName(fileName); err != nil {
		return err
	}

	// check if this user exist
	if !model.FileExist(userName, folderName, fileName) {
		return fmt.Errorf("Error: The File %s doesn't exist.\n", fileName)
	}

	model.DeleteFile(userName, folderName, folderName)

	fmt.Printf("Delete %s in %s / %s successfully.\n", fileName, userName, folderName)

	return nil
}
