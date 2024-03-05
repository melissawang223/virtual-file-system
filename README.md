# Project Title

Create a virtual file system using Golang

## Description

* We could register users.
  * A user name should only contains a-z, A-Z and 0-9.
  * A user name's length should be less than 20
* Each user could have plenty of folders.
  * A folder name should only contains a-z, A-Z and 0-9.
  * A folder name's length should be less than 20
* Each folder could have a optional description
  * The length of the description should less than 50
* Each folder could save plenty of files. 
  * A file name should only contains a-z, A-Z and 0-9.
  * A file name's length should be less than 20
* Each file could have a optional description
    * The length of the description should less than 50

This is the command that support in this virtual file system:\
* User:
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
  * help : Get all the command line
    ```
    help
    ```
  * exit : exit the program
    ```
    exit
    ```

## Getting Started

### Dependencies

* go 1.21+

### Installing

* How/where to download your program
* Any modifications needed to be made to files/folders

### Executing program

* type the following script to start the program

```
go run main.go 
```

## Help

Contact me if you got any trouble


## Authors

Melissa Wang  
melissawang223@gmail.com

