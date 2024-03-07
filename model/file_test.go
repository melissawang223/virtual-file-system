package model

import (
	"reflect"
	"testing"
)

func TestListFile(t *testing.T) {
	initData()
	CreateUser("melissa")
	CreateFolder("melissa", "foldername1", &Folder{
		Name:        "foldername1",
		Description: "d1",
		CreatedAt:   twoHourBefore,
		File:        map[string]*File{},
	})
	CreateFile("melissa", "foldername1", "filename1", &File{
		Name:        "filename1",
		Description: "file1d1",
		CreatedAt:   twoHourBefore,
		FolderName:  "foldername1",
		UserName:    "melissa",
	})
	CreateFile("melissa", "foldername1", "filename2", &File{
		Name:        "filename2",
		Description: "file2d2",
		CreatedAt:   oneHourBefore,
		FolderName:  "foldername1",
		UserName:    "melissa",
	})

	type args struct {
		userName   string
		folderName string
		sortType   string
		sortDir    string
	}
	tests := []struct {
		name string
		args args
		want []File
	}{
		{
			name: "Happy Path: list all files by --sort-name desc",
			args: args{
				userName:   "melissa",
				folderName: "foldername1",
				sortType:   "--sort-name",
				sortDir:    "desc",
			},
			want: []File{
				{
					Name:        "filename2",
					Description: "file2d2",
					CreatedAt:   oneHourBefore,
					FolderName:  "foldername1",
					UserName:    "melissa",
				},
				{
					Name:        "filename1",
					Description: "file1d1",
					CreatedAt:   twoHourBefore,
					FolderName:  "foldername1",
					UserName:    "melissa",
				},
			},
		},
		{
			name: "Happy Path: list all files by --sort-name asc",
			args: args{
				userName:   "melissa",
				folderName: "foldername1",
				sortType:   "--sort-name",
				sortDir:    "asc",
			},
			want: []File{
				{
					Name:        "filename1",
					Description: "file1d1",
					CreatedAt:   twoHourBefore,
					FolderName:  "foldername1",
					UserName:    "melissa",
				},
				{
					Name:        "filename2",
					Description: "file2d2",
					CreatedAt:   oneHourBefore,
					FolderName:  "foldername1",
					UserName:    "melissa",
				},
			},
		},
		{
			name: "Happy Path: list all files by --sort-created desc",
			args: args{
				userName:   "melissa",
				folderName: "foldername1",
				sortType:   "--sort-name",
				sortDir:    "desc",
			},
			want: []File{
				{
					Name:        "filename2",
					Description: "file2d2",
					CreatedAt:   oneHourBefore,
					FolderName:  "foldername1",
					UserName:    "melissa",
				},
				{
					Name:        "filename1",
					Description: "file1d1",
					CreatedAt:   twoHourBefore,
					FolderName:  "foldername1",
					UserName:    "melissa",
				},
			},
		},
		{
			name: "Happy Path: list all files by --sort-created asc",
			args: args{
				userName:   "melissa",
				folderName: "foldername1",
				sortType:   "--sort-name",
				sortDir:    "asc",
			},
			want: []File{
				{
					Name:        "filename1",
					Description: "file1d1",
					CreatedAt:   twoHourBefore,
					FolderName:  "foldername1",
					UserName:    "melissa",
				},
				{
					Name:        "filename2",
					Description: "file2d2",
					CreatedAt:   oneHourBefore,
					FolderName:  "foldername1",
					UserName:    "melissa",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListFile(tt.args.userName, tt.args.folderName, tt.args.sortType, tt.args.sortDir); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListFile() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeleteFile(t *testing.T) {
	initData()
	CreateUser("melissa")
	CreateFolder("melissa", "foldername1", &Folder{
		Name:        "foldername1",
		Description: "d1",
		CreatedAt:   twoHourBefore,
		File:        map[string]*File{},
	})
	CreateFile("melissa", "foldername1", "filename1", &File{
		Name:        "filename1",
		Description: "file1d1",
		CreatedAt:   twoHourBefore,
		FolderName:  "foldername1",
		UserName:    "melissa",
	})
	CreateFile("melissa", "foldername1", "filename2", &File{
		Name:        "filename2",
		Description: "file2d2",
		CreatedAt:   oneHourBefore,
		FolderName:  "foldername1",
		UserName:    "melissa",
	})

	type args struct {
		userName   string
		folderName string
		fileName   string
	}
	type argsForList struct {
		userName   string
		folderName string
		sortType   string
		sortDir    string
	}
	tests := []struct {
		name        string
		args        args
		argsForList argsForList
		want        []File
	}{
		{
			name: "Happy Path: delete file",
			args: args{
				userName:   "melissa",
				folderName: "foldername1",
				fileName:   "filename1",
			},
			argsForList: argsForList{
				userName:   "melissa",
				folderName: "foldername1",
				sortType:   "--sort-name",
				sortDir:    "desc",
			},
			want: []File{
				{
					Name:        "filename2",
					Description: "file2d2",
					CreatedAt:   oneHourBefore,
					FolderName:  "foldername1",
					UserName:    "melissa",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			DeleteFile(tt.args.userName, tt.args.folderName, tt.args.fileName)

			if got := ListFile(tt.argsForList.userName, tt.argsForList.folderName, tt.argsForList.sortType, tt.argsForList.sortDir); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListFolder() = %v, want %v", got, tt.want)
			}
		})
	}
}
