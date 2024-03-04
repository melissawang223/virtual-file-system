package cmd

import (
	"fmt"
	"regexp"
	"virtualFileSystem/folder"
	"virtualFileSystem/user"
)

func CreateUser(args []string) {
	var userName = "melissa"

	if len(args) >= 1 && args[0] != "" {
		userName = args[0]
	} else {
		return
	}

	// create user

	// check invalid char
	usernameConvention := "^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$"
	re, _ := regexp.Compile(usernameConvention)
	if len(userName) > 40 || len(userName) <= 0 || !re.MatchString(userName) {
		fmt.Printf("Error: The %s contain invalid chars.\n", userName)
		return
	}

	// check exist
	if _, ok := user.UsersMap[userName]; ok {
		fmt.Printf("Error: The %s has already existed.\n", userName)
		return
	}

	user.Users = append(user.Users, &user.User{
		Name:    userName,
		Folders: map[string]*folder.Folder{},
	})

	fmt.Printf("Add %s successfully.\n", userName)
}
