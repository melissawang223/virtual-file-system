package cmd

import (
	"fmt"
	"sort"
	"time"
	"virtualFileSystem/folder"
	"virtualFileSystem/user"
)

// list-folders [username] [--sort-name|--sort-created] [asc|desc]

func ListFile(args []string) {
	userName := "melissa"
	sortType := "--sort-name"
	sortDir := "asc"

	if len(args) >= 1 && args[0] != "" {
		userName = args[0]
		if len(args) >= 2 && args[1] != "" {
			sortType = args[1]
		}
		if len(args) >= 3 && args[2] != "" {
			sortDir = args[2]
		}
	} else {
		return
	}

	// check user exist
	if _, ok := user.UsersMap[userName]; !ok {
		_ = fmt.Errorf("Error: The %s doesn't exist.\n", userName)
		return
	}

	// check folder exist
	currentUser := user.UsersMap[userName]
	if len(currentUser.Folders) == 0 {
		fmt.Printf("Warning: The %s doesn't have any folders.\n", userName)
		return
	}

	//list folder
	//[foldername] [description] [created at] [username]
	folders := make([]folder.Folder, 0)
	for _, val := range currentUser.Folders {
		folders = append(folders, *val)
	}

	switch sortType {
	case "--sort-created":
		if sortDir == "des" {
			sort.Slice(folders, func(i, j int) bool {
				return folders[i].CreatedAt < folders[j].CreatedAt
			})
		} else {
			sort.Slice(folders, func(i, j int) bool {
				return folders[i].CreatedAt > folders[j].CreatedAt
			})

		}

	default:
		if sortDir == "des" {
			sort.Slice(folders, func(i, j int) bool {
				return folders[i].Name < folders[j].Name
			})
		} else {
			sort.Slice(folders, func(i, j int) bool {
				return folders[i].Name > folders[j].Name
			})

		}

	}

	for _, val := range folders {
		//[foldername] [description] [created at] [username]
		t := time.Unix(val.CreatedAt, 0)
		fmt.Printf("%s %s %s %s\n", val.Name, val.Description, t.Format(time.RFC3339), userName)
	}
}
