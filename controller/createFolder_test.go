package controller

import "testing"

func TestCreateFolderController(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Happy Path: create folder successfully.",
			args: args{args: []string{"create-folder", "melissa", "foldername", "description"}},
		},
		{
			name: "Happy Path: create folder successfully.",
			args: args{args: []string{"create-folder", "melissa", "foldername"}},
		},
		{
			name: "unHappy Path: register user successfully.",
			args: args{args: []string{"register", "melissa", "foldername!"}},
		},
		{
			name: "unHappy Path: register user successfully.",
			args: args{args: []string{"register", "melissa", "foldername!"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CreateFolderController(tt.args.args)
		})
	}
}
