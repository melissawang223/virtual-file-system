package helper

import (
	"errors"
	"fmt"
	"regexp"
)

var userNameConvention = "^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$"
var folderNameConvention = "^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$"
var fileNameConvention = "^[a-zA-Z0-9]+(?:-[a-zA-Z0-9]+)*$"

func CheckUser(userName string) error {

	// check invalid char
	re, _ := regexp.Compile(userNameConvention)
	if len(userName) > 40 || len(userName) <= 0 || !re.MatchString(userName) {
		fmt.Printf("Error: The username %s contain invalid chars.\n", userName)
		return errors.New("invalid userName")
	}
	return nil
}

func CheckFolder(folderName string) error {

	// check invalid char
	re, _ := regexp.Compile(folderNameConvention)
	if len(folderName) > 40 || len(folderName) <= 0 || !re.MatchString(folderName) {
		fmt.Printf("Error: The folderName %s contain invalid chars.\n", folderName)
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

func CheckFile(fileName string) error {

	// check invalid char
	re, _ := regexp.Compile(fileNameConvention)
	if len(fileName) > 40 || len(fileName) <= 0 || !re.MatchString(fileName) {
		fmt.Printf("Error: The fileName %s contain invalid chars.\n", fileName)
		return errors.New("invalid fileName")
	}
	return nil
}