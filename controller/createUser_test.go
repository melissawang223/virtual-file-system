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
		name string
		args args
	}{
		{
			name: "Happy Path: register user successfully.",
			args: args{args: []string{"register", "melissa"}},
		},
		{
			name: "unHappy Path: register user successfully.",
			args: args{args: []string{"register", "melissa!"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateUserController(tt.args.args)
		})
	}
}
