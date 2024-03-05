package model

import "fmt"

var Users []*User
var UsersMap map[string]*User

type User struct {
	Name    string
	Folders map[string]*Folder
}

func UserExist(userName string) bool {
	if _, ok := UsersMap[userName]; ok {
		fmt.Printf("Error: The %s has already existed.\n", userName)
		return true
	}
	return false
}

func CreateUser(userName string) {
	newUser := &User{
		Name:    userName,
		Folders: map[string]*Folder{},
	}
	Users = append(Users, newUser)
	UsersMap[userName] = newUser
}
