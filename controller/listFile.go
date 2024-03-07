package controller

import (
	"fmt"
	"time"
	"virtualFileSystem/helper"
	"virtualFileSystem/model"
)

func ListFileController(args []string) error {
	userName := ""
	folderName := ""
	sortType := ""
	sortDir := ""

	if len(args) >= 4 {
		userName = args[0]
		folderName = args[1]
		sortType = args[2]
		sortDir = args[3]
		// check sortType and sortDir
		if err := helper.CheckSortTypeAndSortDir(sortType, sortDir); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("Usage: list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]\n")
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
		return fmt.Errorf("Error: The folder %s doesn't exist.\n", folderName)
	}

	// check if folder is empty
	currentUser := model.UsersMap[userName]
	currentFolder := currentUser.Folders[folderName]
	if len(currentFolder.File) == 0 {
		return fmt.Errorf("Warning: The folder %s is empty\n", folderName)
	}

	files := model.ListFile(userName, folderName, sortType, sortDir)

	printFiles(userName, folderName, files)
	return nil
}

func printFiles(userName, folderName string, files []model.File) {
	for _, val := range files {
		t := time.Unix(val.CreatedAt, 0)
		fmt.Printf("%s %s %s %s %s\n", val.Name, val.Description, t.Format(time.DateTime), folderName, userName)
	}
}
