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

var storage = "local.txt"

func main() {

	readToMemory()

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("Catch")
		saveMemory()
		os.Exit(1)
	}()
	printHelp()
out:
	for {
		buf := bufio.NewReader(os.Stdin)

		fmt.Print("> ")
		sentence, err := buf.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		} else {

			args := strings.Fields(sentence)
			if len(args) >= 1 && args[0] != "" {
				switch args[0] {
				case "register":
					cmd.CreateUser(args[1:])

				case "create-folder":
					cmd.CreateFolder(args[1:])
				case "delete-folder":
					cmd.DeleteFolder(args[1:])
				case "list-folders":
					cmd.ListFolder(args[1:])
				case "rename-folder":
					cmd.ListFolder(args[1:])

				case "create-file":
					cmd.CreateFolder(args[1:])
				case "delete-file":
					cmd.DeleteFolder(args[1:])
				case "list-files":
					cmd.ListFolder(args[1:])

				case "help":
					printHelp()
				case "exit":
					saveMemory()
					break out
				default:
					fmt.Printf("Your input: %s did not match any operation we support.\n", args[0])
					fmt.Println("Please check out the supporting operation with 'help', or type 'exit' to leave")
				}
			}

		}

	}
}

func readToMemory() {
	user.Users = make([]*user.User, 0)
	user.UsersMap = map[string]*user.User{}

	byt, err := os.ReadFile(storage)
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

}

func saveMemory() {

	file, _ := json.MarshalIndent(user.Users, "", " ")

	err := os.WriteFile(storage, file, 0644)
	if err != nil {
		log.Fatal(err)
	}

}

func printHelp() {
	fmt.Println("> User:")
	fmt.Println(">\t register [username]")
	fmt.Println("> Folder:")
	fmt.Println(">\t create-folder [username] [foldername] [description]?")
	fmt.Println(">\t delete-folder [username] [foldername]")
	fmt.Println(">\t list-folders [username] [--sort-name|--sort-created] [asc|desc]")
	fmt.Println(">\t rename-folder [username] [foldername] [new-folder-name]")
	fmt.Println("> File:")
	fmt.Println(">\t create-file [username] [foldername] [filename] [description]?")
	fmt.Println(">\t delete-file [username] [foldername] [filename]")
	fmt.Println(">\t list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]")
	fmt.Println("> Others:")
	fmt.Println(">\t help")
	fmt.Println(">\t exit")
}
