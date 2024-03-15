package controller

import (
	"testing"
	"virtualFileSystem/model"
)

func InitData() {
	model.Users = make([]*model.User, 0)
	model.UsersMap = map[string]*model.User{}
	model.FolderMap = map[[2]string]*model.Folder{}
	model.FileMap = map[[3]string]*model.File{}
}

func TestRenameFolderController(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Happy Path: rename folder successfully.",
			args:    args{args: []string{"melissa", "foldername", "newfoldername"}},
			wantErr: false,
		},
		{
			name:    "unHappy Path: rename folder failed: username contains special char",
			args:    args{args: []string{"melissa!", "foldername!", "foldername2"}},
			wantErr: true,
		},
		{
			name:    "unHappy Path: rename folder failed: foldername contains special char",
			args:    args{args: []string{"melissa", "foldername", "foldername!"}},
			wantErr: true,
		},
		{
			name:    "unHappy Path: username does not exist",
			args:    args{args: []string{"melissa2", "foldername", "foldername3"}},
			wantErr: true,
		},
		{
			name:    "unHappy Path: foldername does not exist",
			args:    args{args: []string{"melissa", "notexistfoldername"}},
			wantErr: true,
		},
		{
			name:    "unHappy Path: rename folder failed: foldername contains special char",
			args:    args{args: []string{"melissa", "foldername", "existfolder"}},
			wantErr: true,
		},
	}
	InitData()
	CreateUserController([]string{"melissa"})
	CreateFolderController([]string{"melissa", "foldername"})
	CreateFolderController([]string{"melissa", "existfolder"})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RenameFolderController(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("RenameFolderController() error = %v, wantErr %v", err, tt.wantErr)
			}
			// compare file...
		})
	}
}
