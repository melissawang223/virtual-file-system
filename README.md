# Project Title

Create a virtual file system using Golang

## Description

We will storage the data into local.txt, and read from it when we start the program.

### User: We could register users.
  * A username should only contain a-z, A-Z and 0-9.
  * A username's length should be less than 10
### Folder: Each user could have plenty of folders.
  * A folder name should only contain a-z, A-Z and 0-9.
  * A folder name's length should be less than 15
  * Each folder could have a optional description
    * The length of the description should less than 30
### File: Each folder could save plenty of files. 
  * A file name should only contain a-z, A-Z and 0-9.
  * A file name's length should be less than 20
  * Each file could have a optional description
      * The length of the description should less than 30
### Exit: Leave the program by ctrl+c or type exit
    * the data will be store to local.txt

### Help: List down all the supported function
 

## RoadMap
This is the command that support in this virtual file system:
### User:
* register username
  * check the username is valid 
    * username is not empty
    * length < 10 
    * only a-z, A-Z and 0-1
  * check the username has not register before
```
    register [username]
```
  
* Folder:
  * create folder: 
    ```
    create-folder [username] [foldername] [description]?
    ```
  * delete-folder
    ```
    delete-folder [username] [foldername]
    ```
  * list-folders
    ```
    list-folders [username] [--sort-name|--sort-created] [asc|desc]
    ```
  * rename-folder
    ```
    rename-folder [username] [foldername] [new-folder-name]
    ```
  * File:
    * create-file
    ```
    create-file [username] [foldername] [filename] [description]?
    ```
    * delete-file
    ```
    delete-file [username] [foldername] [filename]
    ```
  * list-files
    ```
     list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]
    ```


* Others:
  * help : list all the supported command line
    ```
    help
    ```
  * exit : Exit the program. We will save the data into the local.txt file
    ```
    exit
    ```

## Getting Started

### Dependencies

* go 1.21+

### Executing program

At the very first time, you may need to run 'make inti' to create a local storage
```
make init
```


* Then we could start the program by 'make run' or 'go run main.go'

```
make run
```
or 
```
go run main.go 
```

## Help

Contact me if you got any trouble


## Authors

Melissa Wang  
melissawang223@gmail.com

