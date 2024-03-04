package folder

import "virtualFileSystem/file"

type Folder struct {
	Name        string
	Description string
	CreatedAt   int64
	File        map[string]*file.File
}

/*
sort.Slice(folders, func(i, j int) bool {
	return users[i].Name < users[j].Name
})


sort.Slice(folders, func(i, j int) bool {
	return users[i].CreatedAt < users[j].CreatedAt
})


*/
