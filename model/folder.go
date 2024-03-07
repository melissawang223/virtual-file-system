package model

import (
	"sort"
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
		return true
	}
	return false
}

func DeleteFolder(userName, folderName string) {

	//delete folder
	currentFolder := UsersMap[userName].Folders

	// delete files inside the folder
	if currentFolder[folderName].File != nil {
		for _, val := range currentFolder[folderName].File {
			fileName := val.Name
			DeleteFile(userName, folderName, fileName)
		}
	}

	delete(FolderMap, [2]string{userName, folderName})
	delete(currentFolder, folderName)
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

func ListFolder(userName, sortType, sortDir string) []Folder {

	currentUser := UsersMap[userName]
	folders := make([]Folder, 0)
	for _, val := range currentUser.Folders {
		folders = append(folders, *val)
	}

	switch sortType {
	case "--sort-created":
		if sortDir == "desc" {
			sort.Slice(folders, func(i, j int) bool {
				return folders[i].CreatedAt > folders[j].CreatedAt
			})
		} else {
			sort.Slice(folders, func(i, j int) bool {
				return folders[i].CreatedAt < folders[j].CreatedAt
			})

		}

	default:
		if sortDir == "desc" {
			sort.Slice(folders, func(i, j int) bool {
				return folders[i].Name > folders[j].Name
			})
		} else {
			sort.Slice(folders, func(i, j int) bool {

				return folders[i].Name < folders[j].Name
			})

		}

	}

	return folders
}
