# Project Title

Create a virtual file system using Golang

## Description

We could register users. 
Each user could have plenty of folders.
Each folder could save plenty of Files

This is the command that support in this virtual file system
User:
    register [username]
Folder:
    create-folder [username] [foldername] [description]?
    delete-folder [username] [foldername]
    list-folders [username] [--sort-name|--sort-created] [asc|desc]
    rename-folder [username] [foldername] [new-folder-name]
File: 
    create-file [username] [foldername] [filename] [description]?
    delete-file [username] [foldername] [filename]
    list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]
Others:
    help
    exit


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
```

```

## Authors

Melissa Wang  
melissawang223@gmail.com

