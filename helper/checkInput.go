package helper

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"regexp"
	"strings"
)

var nameConvention = `"([a-zA-Z0-9\s]+)"|^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$`

func CheckUserName(userName string) error {

	// check invalid char
	re, _ := regexp.Compile(nameConvention)
	if len(userName) > 10 || len(userName) <= 0 || !re.MatchString(userName) {
		return fmt.Errorf("Error: The username %s contain invalid chars.( Or Length is over 40 )\n", userName)
	}
	return nil
}

func CheckFolderName(folderName string) error {

	// check invalid char
	re, _ := regexp.Compile(nameConvention)
	if len(folderName) > 15 || len(folderName) <= 0 || !re.MatchString(folderName) {
		return fmt.Errorf("Error: The foldername %s contain invalid chars.\n", folderName)
	}
	return nil
}

func CheckFileDescription(fileDes string) error {

	// check invalid char
	if len(fileDes) >= 30 {
		fmt.Printf("Error: The Description %s has length over 30.\n", fileDes)
		return errors.New("invalid Description")
	}
	return nil
}

func CheckFileName(fileName string) error {

	// check invalid char
	re, _ := regexp.Compile(nameConvention)
	if len(fileName) > 20 || len(fileName) <= 0 || !re.MatchString(fileName) {
		return fmt.Errorf("Error: The fileName %s contain invalid chars.\n", fileName)
	}
	return nil
}

func CheckSortTypeAndSortDir(sortType, sortDir string) error {
	if sortType != "--sort-name" && sortType != "--sort-created" {
		return fmt.Errorf("Usage: list-folders [username] [--sort-name|--sort-created] [asc|desc]`")
	}

	if sortDir != "asc" && sortDir != "desc" {
		return fmt.Errorf("Usage: list-folders [username] [--sort-name|--sort-created] [asc|desc]`")
	}
	return nil
}

func CheckArguments(input string) ([]string, error) {
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
