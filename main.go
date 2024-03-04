package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	vfs "virtualFileSystem/cmd"
	"virtualFileSystem/user"
)

func readToMemory() {
	user.Users = make([]user.User, 0)

	//User
	byt, err := os.ReadFile("localUser.txt")
	if err != nil {
		fmt.Println("No local user storage")
	}

	if err := json.Unmarshal(byt, &user.Users); err != nil {
		panic(err)
	}

	fmt.Println(user.Users)

	// Folder

	// File

}

func main() {

	defer func() {
		saveMemory()
	}()
	readToMemory()
	/*read, err := os.ReadFile("localUser.txt")
	if err != nil {
		fmt.Println("No local storage")
	}

	_, err = os.Stdout.Write(read)
	if err != nil {
		return
	}

	// `b` contains everything your file has.
	// This writes it to the Standard Out.
	temp := []byte("Here is a string3....")
	_, err = os.Stdout.Write(temp)
	if err != nil {
		return
	}
	*/

	vfs.Execute()

}

func saveMemory() {
	// You can also write it to a file as a whole.

	file, _ := json.MarshalIndent(user.Users, "", " ")

	err := os.WriteFile("localUser.txt", file, 0644)
	if err != nil {
		log.Fatal(err)
	}

}
