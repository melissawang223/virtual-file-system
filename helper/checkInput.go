package helper

import (
	"errors"
	"fmt"
	"regexp"
)

var userNameConvention = "^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$"
var folderNameConvention = "^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$"
var fileNameConvention = "^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$"

func CheckUserName(userName string) error {

	// check invalid char
	re, _ := regexp.Compile(userNameConvention)
	if len(userName) > 10 || len(userName) <= 0 || !re.MatchString(userName) {
		errMsg := fmt.Sprintf("Error: The username %s contain invalid chars.( Or Length is over 40 )\n", userName)
		return errors.New(errMsg)
	}
	return nil
}

func CheckFolderName(folderName string) error {

	// check invalid char
	re, _ := regexp.Compile(folderNameConvention)
	if len(folderName) > 15 || len(folderName) <= 0 || !re.MatchString(folderName) {
		fmt.Printf("Error: The foldername %s contain invalid chars.\n", folderName)
		return errors.New("invalid folderName")
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
		fmt.Printf("Error: The fileName %s contain invalid chars.\n", fileName)
		return errors.New("invalid fileName")
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
