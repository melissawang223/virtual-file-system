package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"virtualFileSystem/cmd"
	"virtualFileSystem/user"
)

func readToMemory() {
	user.Users = make([]*user.User, 0)
	user.UsersMap = map[string]*user.User{}

	//User
	byt, err := os.ReadFile("localUser.txt")
	if err != nil {
		fmt.Println("No local user storage")
	}

	if err := json.Unmarshal(byt, &user.Users); err != nil {
		panic(err)
	}

	fmt.Println(user.Users)

	for _, val := range user.Users {
		user.UsersMap[val.Name] = val
	}

	// Folder

	// File

}

func main() {
	defer func() {
		fmt.Println("In Defer")
		saveMemory()
	}()

	readToMemory()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		saveMemory()
		os.Exit(1)
	}()

	for {
		buf := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		sentence, err := buf.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		} else {

			//vfs.Execute()
			args := strings.Fields(sentence)
			if len(args) >= 1 && args[0] != "" {
				switch args[0] {
				case "register":
					cmd.CreateUser(args[1:])
				case "create-folder":
					cmd.CreateFolder(args[1:])
				case "delete-folder":
					cmd.DeleteFolder(args[1:])
				case "list-folder":
					cmd.ListFolder(args[1:])
				}
			}

		}

	}
}

func saveMemory() {

	// user
	file, _ := json.MarshalIndent(user.Users, "", " ")

	err := os.WriteFile("localUser.txt", file, 0644)
	if err != nil {
		log.Fatal(err)
	}

}
