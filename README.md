# Virtual File System

Create a virtual file system using Golang

* main branch is the basic function
* feature/bonus branch is the bonus function

## Description

We will store the data into local.txt, and read from it when we start the program.

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
#### register username
* username check
  * username is not empty
  * username length < 10 
  * only a-z, A-Z and 0-9
* check the username has not registered before
```
    register [username]
```

### Folder:
#### create folder 
* username check
  * username is not empty
  * username length < 10
  * only a-z, A-Z and 0-9
* check the user exist
* foldername check
  * foldername is not empty
  * foldername length < 15
  * only a-z, A-Z and 0-9
* check the foldername has not existed
```
  create-folder [username] [foldername] [description]?
 ```
#### delete-folder
* username check
  * username is not empty
  * username length < 10
  * only a-z, A-Z and 0-9
* check the user exist
* foldername check
  * foldername is not empty
  * foldername length < 15
  * only a-z, A-Z and 0-9
* check the foldername exist 
  ```
  delete-folder [username] [foldername]
  ```
#### list-folders
* username check
  * username is not empty
  * username length < 10
  * only a-z, A-Z and 0-9
* check the user exist
* check sortType(--sort-name and --sort-created) 
* check sorDirection (asc and desc)
* check users does not have empty folder
```
  list-folders [username] [--sort-name|--sort-created] [asc|desc]
```
#### rename-folder
* username check
  * username is not empty
  * username length < 10
  * only a-z, A-Z and 0-9
* foldername check
  * foldername is not empty
  * foldername length < 15
  * only a-z, A-Z and 0-9
  * check the foldername exist
* newfoldername check
  * newfoldername is not empty
  * newfoldername length < 15
  * only a-z, A-Z and 0-9
* check the newfoldername has not exist
```
  rename-folder [username] [foldername] [new-folder-name]
```
### File:
#### create-file
* username check
  * username is not empty
  * username length < 10
  * only a-z, A-Z and 0-9
* check the user exist
* foldername check
  * foldername is not empty
  * foldername length < 15
  * only a-z, A-Z and 0-9
* check the foldername exist
* filename check
  * filename is not empty
  * filename length < 20
  * only a-z, A-Z and 0-9
* check the filename has not existed
```
  create-file [username] [foldername] [filename] [description]?
```
#### delete-file
* username check
  * username is not empty
  * username length < 10
  * only a-z, A-Z and 0-9
* check the user exist
* foldername check
  * foldername is not empty
  * foldername length < 15
  * only a-z, A-Z and 0-9
* check the foldername exist
* filename check
  * filename is not empty
  * filename length < 20
  * only a-z, A-Z and 0-9
* check the filename has existed
```
  delete-file [username] [foldername] [filename]
```
#### list-files
* username check
  * username is not empty
  * username length < 10
  * only a-z, A-Z and 0-9
* check the user exist
* foldername check
  * foldername is not empty
  * foldername length < 15
  * only a-z, A-Z and 0-9
* check the foldername exist
* check sortType(--sort-name and --sort-created)
* check sorDirection (asc and desc)
* check the folder does not have empty files
```
  list-files [username] [foldername] [--sort-name|--sort-created] [asc|desc]
```


### Others:
#### help : list all the supported command line
```
  help
```
#### exit : Exit the program. We will save the data into the local.txt file
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


* Then we could start the program by building the file
```
make build
```

* Then we could run the program with 
```
make run
```

### Run the unit tests

* Run all the unit tests
```
  make test
```
* Run all test coverage
```
  make test_coverage
```

## Help

Contact me if you got any trouble on running the program


## Authors

Melissa Wang  
melissawang223@gmail.com

