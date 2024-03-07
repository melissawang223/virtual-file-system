package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
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
		saveMemory()
		os.Exit(1)
	}()
out:
	for {
		buf := bufio.NewReader(os.Stdin)

		fmt.Print("> ")
		sentence, err := buf.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		} else {

			args, argErr := processArg(sentence)
			if argErr != nil && len(args) >= 1 && args[0] != "" {
				switch args[0] {
				case "register":
					fmt.Fprintf(os.Stderr, "Usage: register [username]\n")

				case "create-folder":
					fmt.Fprintf(os.Stderr, "Usage: create-folder [username] [foldername] [description]?`\n")

				case "delete-folder":
					fmt.Fprintf(os.Stderr, "Usage: delete-folder [username] [foldername]\n")

				case "list-folders":
					fmt.Fprintf(os.Stderr, "Usage: list-folders [username] [--sort-name|--sort-created] [asc|desc]\n")

				case "rename-folder":
					fmt.Fprintf(os.Stderr, "Usage: rename-folder [username] [foldername] [new-folder-name]\n")

				case "create-file":
					fmt.Fprintf(os.Stderr, "Usage: create-file [username] [foldername] [filename] [description]?\n")

				case "delete-file":
					fmt.Fprintf(os.Stderr, "Usage: `delete-file [username] [foldername] [filename]\n")

				case "list-files":
					fmt.Fprintf(os.Stderr, "Usage: list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]\n")

				default:
					fmt.Println(" Error: Unrecognized command")
					fmt.Printf("Your input: %s did not match any operation we support.\n", args[0])
					fmt.Println("Please check out the supporting operation with 'help', or type 'exit' to leave")
				}

			} else if len(args) >= 1 && args[0] != "" {

				switch args[0] {
				case "register":
					err = controller.CreateUserController(args[1:])
					if err != nil {
						fmt.Fprintf(os.Stderr, err.Error())
					}

				case "create-folder":
					err = controller.CreateFolderController(args[1:])
					if err != nil {
						fmt.Fprintf(os.Stderr, err.Error())
					}
				case "delete-folder":
					err = controller.DeleteFolderController(args[1:])
					if err != nil {
						fmt.Fprintf(os.Stderr, err.Error())
					}
				case "list-folders":
					err = controller.ListFolderController(args[1:])
					if err != nil {
						fmt.Fprintf(os.Stderr, err.Error())
					}
				case "rename-folder":
					err = controller.RenameFolderController(args[1:])
					if err != nil {
						fmt.Fprintf(os.Stderr, err.Error())
					}

				case "create-file":
					err = controller.CreateFileController(args[1:])
					if err != nil {
						fmt.Fprintf(os.Stderr, err.Error())
					}
				case "delete-file":
					err = controller.DeleteFileController(args[1:])
					if err != nil {
						fmt.Fprintf(os.Stderr, err.Error())
					}
				case "list-files":
					err = controller.ListFileController(args[1:])
					if err != nil {
						fmt.Fprintf(os.Stderr, err.Error())
					}

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

			} else {
				fmt.Println(" Error: Unrecognized command")
				fmt.Printf("Your input: %s did not match any operation we support.\n", args[0])
				fmt.Println("Please check out the supporting operation with 'help', or type 'exit' to leave")
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

func processArg(input string) ([]string, error) {
	args := make([]string, 0)

	temp := ""

	hsMap := map[string]string{}

	for i := 0; i < len(input); i++ {
		cur := input[i]
		if cur == '"' {
			pre := i
			i++
			for i < len(input) && input[i] != '"' {
				i++
			}
			if i < len(input) && input[i] == '"' {
				key := uuid.New().String()
				hsMap[key] = input[pre : i+1]
				temp += key
			} else {
				return []string{}, fmt.Errorf("Invalid input")
			}
		} else {
			temp += string(cur)
		}
	}

	args = strings.Fields(temp)

	for idx, val := range args {
		if double, ok := hsMap[val]; ok {
			args[idx] = double
		}
	}
	return args, nil
}
