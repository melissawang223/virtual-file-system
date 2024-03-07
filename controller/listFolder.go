package controller

import (
	"fmt"
	"time"
	"virtualFileSystem/helper"
	"virtualFileSystem/model"
)

func ListFolderController(args []string) error {
	userName := ""
	sortType := ""
	sortDir := ""

	if len(args) >= 3 {
		userName = args[0]
		sortType = args[1]
		sortDir = args[2]
		// check sortType and sortDir
		if err := helper.CheckSortTypeAndSortDir(sortType, sortDir); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("Usage: list-folders [username] [--sort-name|--sort-created] [asc|desc]\n")
	}

	// check userName is valid or not
	if err := helper.CheckUserName(userName); err != nil {
		return err
	}

	// check if this user exist
	if !model.UserExist(userName) {
		return fmt.Errorf("Error: The %s has not existed.\n", userName)
	}

	// check folder exist
	currentUser := model.UsersMap[userName]
	if len(currentUser.Folders) == 0 {
		return fmt.Errorf("Warning: The %s doesn't have any folders.\n", userName)
	}

	//list folder
	folders := model.ListFolder(userName, sortType, sortDir)
	printFolders(userName, folders)
	return nil
}

func printFolders(userName string, folders []model.Folder) {
	for _, val := range folders {
		t := time.Unix(val.CreatedAt, 0)
		fmt.Printf("%s %s %s %s\n", val.Name, val.Description, t.Format(time.DateTime), userName)
	}
}
