package helper

import (
	"errors"
	"fmt"
	"regexp"
)

var userNameConvention = `"([a-zA-Z0-9\s]+)"|([a-zA-Z0-9]+)`
var folderNameConvention = `"([a-zA-Z0-9\s]+)"|([a-zA-Z0-9]+)`
var fileNameConvention = `"([a-zA-Z0-9\s]+)"|([a-zA-Z0-9]+)`

func CheckUserName(userName string) error {

	// check invalid char
	re, _ := regexp.Compile(userNameConvention)
	if len(userName) > 10 || len(userName) <= 0 || !re.MatchString(userName) {
		return fmt.Errorf("Error: The username %s contain invalid chars.( Or Length is over 40 )\n", userName)
	}
	return nil
}

func CheckFolderName(folderName string) error {

	// check invalid char
	re, _ := regexp.Compile(folderNameConvention)
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
	re, _ := regexp.Compile(fileNameConvention)
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
