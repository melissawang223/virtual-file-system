package user

import "virtualFileSystem/folder"

var Users []*User
var UsersMap map[string]*User

type User struct {
	Name    string
	Folders map[string]*folder.Folder
}
