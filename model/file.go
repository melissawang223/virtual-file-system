package model

import (
	"fmt"
	"sort"
	"time"
)

var FileMap map[[3]string]*File

type File struct {
	Name        string
	Description string
	CreatedAt   int64

	FolderName string
	UserName   string
}

func CreateFile(userName, folderName, fileName string, file *File) {
	currentUser := UsersMap[userName]
	currentFolder := currentUser.Folders[folderName]
	currentFolder.File[fileName] = file

	FileMap[[3]string{userName, folderName, fileName}] = file
}

func FileExist(userName, folderName, fileName string) bool {
	if _, ok := FileMap[[3]string{userName, folderName, fileName}]; ok {
		fmt.Printf("Error: The %s has already existed.\n", fileName)
		return true
	}
	return false
}

func DeleteFile(userName, folderName, fileName string) {

	//delete folder
	currentUser := UsersMap[userName]
	currentFolder := currentUser.Folders[folderName]
	delete(currentFolder.File, folderName)
	delete(FileMap, [3]string{userName, folderName, fileName})
}

// `[filename]` `[description]` `[created at]` `[foldername]` `[username]`
func ListFile(userName, folderName, sortType, sortDir string) {

	//[foldername] [description] [created at] [username]
	currentUser := UsersMap[userName]
	currentFolder := currentUser.Folders[folderName]
	files := make([]File, 0)
	for _, val := range currentFolder.File {
		files = append(files, *val)
	}

	switch sortType {
	case "--sort-created":
		if sortDir == "des" {
			sort.Slice(files, func(i, j int) bool {
				return files[i].CreatedAt < files[j].CreatedAt
			})
		} else {
			sort.Slice(files, func(i, j int) bool {
				return files[i].CreatedAt > files[j].CreatedAt
			})

		}

	default:
		if sortDir == "des" {
			sort.Slice(files, func(i, j int) bool {
				return files[i].Name < files[j].Name
			})
		} else {
			sort.Slice(files, func(i, j int) bool {
				return files[i].Name > files[j].Name
			})

		}

	}

	for _, val := range files {
		t := time.Unix(val.CreatedAt, 0)
		fmt.Printf("%s %s %s %s %s\n", val.Name, val.Description, t.Format(time.RFC3339), folderName, userName)
	}
}
