package controller

import (
	"testing"
	"virtualFileSystem/model"
)

func init() {
	model.Users = make([]*model.User, 0)
	model.UsersMap = map[string]*model.User{}
	model.FolderMap = map[[2]string]*model.Folder{}
	model.FileMap = map[[3]string]*model.File{}
}

func TestCreateFolderController(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Happy Path: create folder successfully.",
			args:    args{args: []string{"melissa", "foldername", "description"}},
			wantErr: false,
		},
		{
			name:    "Happy Path: create folder without description successfully.",
			args:    args{args: []string{"melissa", "foldername1"}},
			wantErr: false,
		},
		{
			name:    "unHappy Path: create folder failed: foldername contains special char",
			args:    args{args: []string{"melissa", "foldername!"}},
			wantErr: true,
		},
		{
			name:    "unHappy Path: username does not exist",
			args:    args{args: []string{"melissa2", "foldername!"}},
			wantErr: true,
		},
	}

	CreateUserController([]string{"melissa"})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateFolderController(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("CreateFolderController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
