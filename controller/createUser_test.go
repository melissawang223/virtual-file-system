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

func TestCreateUserController(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Happy Path: register user successfully.",
			args:    args{args: []string{"melissa"}},
			wantErr: false,
		},
		{
			name:    "unHappy Path: username contains special char.",
			args:    args{args: []string{"melissa!"}},
			wantErr: true,
		},
		{
			name:    "unHappy Path: username is empty",
			args:    args{args: []string{""}},
			wantErr: true,
		},
		{
			name:    "unHappy Path: user has already existed.",
			args:    args{args: []string{"melissa"}},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateUserController(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("CreateUserController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
