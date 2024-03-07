package model

import (
	"reflect"
	"testing"
	"time"
)

var oneHourBefore int64
var twoHourBefore int64

func initData() {
	Users = make([]*User, 0)
	UsersMap = map[string]*User{}
	FolderMap = map[[2]string]*Folder{}
	FileMap = map[[3]string]*File{}

	oneHourBefore = time.Now().Add(-1 * time.Hour).Unix()
	twoHourBefore = time.Now().Add(-2 * time.Hour).Unix()

}

func TestFolderExist(t *testing.T) {
	initData()
	CreateUser("melissa")
	CreateFolder("melissa", "foldername1", &Folder{
		Name:        "foldername1",
		Description: "d1",
		CreatedAt:   twoHourBefore,
		File:        map[string]*File{},
	})

	type args struct {
		userName   string
		folderName string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Happy Path: Check foldername1 exist",
			args: args{
				userName:   "melissa",
				folderName: "foldername1",
			},
			want: true,
		},
		{
			name: "UnHappy Path: Check foldername2 not exist",
			args: args{
				userName:   "melissa",
				folderName: "foldername2",
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FolderExist(tt.args.userName, tt.args.folderName); got != tt.want {
				t.Errorf("FolderExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestListFolder(t *testing.T) {
	initData()
	CreateUser("melissa")
	CreateFolder("melissa", "foldername1", &Folder{
		Name:        "foldername1",
		Description: "d1",
		CreatedAt:   twoHourBefore,
		File:        map[string]*File{},
	})
	CreateFolder("melissa", "foldername2", &Folder{
		Name:        "foldername2",
		Description: "d2",
		CreatedAt:   oneHourBefore,
		File:        map[string]*File{},
	})

	type args struct {
		userName string
		sortType string
		sortDir  string
	}
	tests := []struct {
		name string
		args args
		want []Folder
	}{
		{
			name: "Happy Path: list all folders by --sort-name asc ",
			args: args{
				userName: "melissa",
				sortType: "--sort-name",
				sortDir:  "asc",
			},
			want: []Folder{
				{
					Name:        "foldername1",
					Description: "d1",
					CreatedAt:   twoHourBefore,
					File:        map[string]*File{},
				},
				{
					Name:        "foldername2",
					Description: "d2",
					CreatedAt:   oneHourBefore,
					File:        map[string]*File{},
				},
			},
		},
		{
			name: "Happy Path: list all folders by --sort-name desc ",
			args: args{
				userName: "melissa",
				sortType: "--sort-name",
				sortDir:  "desc",
			},
			want: []Folder{
				{
					Name:        "foldername2",
					Description: "d2",
					CreatedAt:   oneHourBefore,
					File:        map[string]*File{},
				},
				{
					Name:        "foldername1",
					Description: "d1",
					CreatedAt:   twoHourBefore,
					File:        map[string]*File{},
				},
			},
		},
		{
			name: "Happy Path: list all folders by --sort-created asc ",
			args: args{
				userName: "melissa",
				sortType: "--sort-created",
				sortDir:  "asc",
			},
			want: []Folder{
				{
					Name:        "foldername1",
					Description: "d1",
					CreatedAt:   twoHourBefore,
					File:        map[string]*File{},
				},
				{
					Name:        "foldername2",
					Description: "d2",
					CreatedAt:   oneHourBefore,
					File:        map[string]*File{},
				},
			},
		},
		{
			name: "Happy Path: list all folders by --sort-created desc ",
			args: args{
				userName: "melissa",
				sortType: "--sort-created",
				sortDir:  "desc",
			},
			want: []Folder{
				{
					Name:        "foldername2",
					Description: "d2",
					CreatedAt:   oneHourBefore,
					File:        map[string]*File{},
				},
				{
					Name:        "foldername1",
					Description: "d1",
					CreatedAt:   twoHourBefore,
					File:        map[string]*File{},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ListFolder(tt.args.userName, tt.args.sortType, tt.args.sortDir); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListFolder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReNameFolder(t *testing.T) {
	initData()

	type args struct {
		userName      string
		folderName    string
		newfolderName string
	}
	type argsForList struct {
		userName string
		sortType string
		sortDir  string
	}
	tests := []struct {
		name        string
		args        args
		argsForList argsForList
		want        []Folder
	}{
		{
			name: "Happy Path: rename foldername1 to foldername3 ",
			args: args{
				userName:      "melissa",
				folderName:    "foldername1",
				newfolderName: "foldername3",
			},
			argsForList: argsForList{
				userName: "melissa",
				sortType: "--sort-name",
				sortDir:  "asc",
			},
			want: []Folder{
				{
					Name:        "foldername2",
					Description: "d2",
					CreatedAt:   oneHourBefore,
					File:        map[string]*File{},
				},
				{
					Name:        "foldername3",
					Description: "d1",
					CreatedAt:   twoHourBefore,
					File:        map[string]*File{},
				},
			},
		},
	}

	CreateUser("melissa")
	CreateFolder("melissa", "foldername1", &Folder{
		Name:        "foldername1",
		Description: "d1",
		CreatedAt:   twoHourBefore,
		File:        map[string]*File{},
	})
	CreateFolder("melissa", "foldername2", &Folder{
		Name:        "foldername2",
		Description: "d2",
		CreatedAt:   oneHourBefore,
		File:        map[string]*File{},
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReNameFolder(tt.args.userName, tt.args.folderName, tt.args.newfolderName)

			if got := ListFolder(tt.argsForList.userName, tt.argsForList.sortType, tt.argsForList.sortDir); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListFolder() = %v, want %v", got, tt.want)
			}
		})
	}
}
