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
	"virtualFileSystem/controller"
	"virtualFileSystem/model"
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
					controller.CreateUserController(args[1:])

				case "create-folder":
					controller.CreateFolderController(args[1:])
				case "delete-folder":
					controller.DeleteFolderController(args[1:])
				case "list-folders":
					controller.ListFolderController(args[1:])
				case "rename-folder":
					controller.ListFolderController(args[1:])

				case "create-file":
					controller.CreateFileController(args[1:])
				case "delete-file":
					controller.DeleteFileController(args[1:])
				case "list-files":
					controller.ListFileController(args[1:])

				case "help":
					printHelp()
				case "exit":
					saveMemory()
					break out
				default:
					fmt.Println(" Error: Unrecognized command")
					fmt.Printf("Your input: %s did not match any operation we support.\n", args[0])
					fmt.Println("Please check out the supporting operation with 'help', or type 'exit' to leave")
				}
			}

		}

	}
}

func init() {
	model.Users = make([]*model.User, 0)
	model.UsersMap = map[string]*model.User{}
	model.FolderMap = map[[2]string]*model.Folder{}
	model.FileMap = map[[3]string]*model.File{}
}

func readToMemory() {

	byt, err := os.ReadFile(storage)
	if err != nil {
		fmt.Println("No local user storage")
	}

	if err := json.Unmarshal(byt, &model.Users); err != nil {
		panic(err)
	}

	for _, user := range model.Users {
		model.UsersMap[user.Name] = user
		for _, folder := range user.Folders {
			model.FolderMap[[2]string{user.Name, folder.Name}] = folder
			for _, file := range folder.File {
				model.FileMap[[3]string{user.Name, folder.Name, file.Name}] = file
			}
		}
	}

}

func saveMemory() {

	file, _ := json.MarshalIndent(model.Users, "", " ")

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
