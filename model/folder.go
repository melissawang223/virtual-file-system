package model

import (
	"fmt"
	"sort"
	"time"
)

var FolderMap map[[2]string]*Folder

type Folder struct {
	Name        string
	Description string
	CreatedAt   int64
	File        map[string]*File
}

func CreateFolder(userName, folderName string, newFolder *Folder) {
	currentUser := UsersMap[userName]
	currentUser.Folders[folderName] = newFolder

	FolderMap[[2]string{userName, folderName}] = newFolder
}

func FolderExist(userName, folderName string) bool {
	if _, ok := FolderMap[[2]string{userName, folderName}]; ok {
		fmt.Printf("Error: The %s has already existed.\n", userName)
		return true
	}
	return false
}

func DeleteFolder(userName, folderName string) {

	//delete folder
	currentFolder := UsersMap[userName].Folders
	delete(currentFolder, folderName)

	for _, val := range currentFolder[folderName].File {
		fileName := val.Name
		DeleteFile(userName, folderName, fileName)
	}

	delete(FolderMap, [2]string{userName, folderName})
}

func ReNameFolder(userName, folderName, newfolderName string) {

	//rename folder
	oldFolder := UsersMap[userName].Folders[folderName]
	newFolder := &Folder{
		Name:        newfolderName,
		Description: oldFolder.Description,
		CreatedAt:   oldFolder.CreatedAt,
		File:        oldFolder.File,
	}

	DeleteFolder(userName, folderName)
	CreateFolder(userName, newfolderName, newFolder)
}

// list folder
func ListFolder(userName, sortType, sortDir string) {

	//[foldername] [description] [created at] [username]
	currentUser := UsersMap[userName]
	folders := make([]Folder, 0)
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
