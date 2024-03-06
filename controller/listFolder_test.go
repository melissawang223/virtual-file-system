package controller

import "testing"

func TestListFolderController(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Happy Path: list-folder --sort-name asc successfully.",
			args:    args{args: []string{"melissa", "--sort-name", "asc"}},
			wantErr: false,
		},
		{
			name:    "Happy Path: list-folder --sort-created asc successfully.",
			args:    args{args: []string{"melissa", "--sort-created", "asc"}},
			wantErr: false,
		},
		{
			name:    "unHappy Path: sort is not support",
			args:    args{args: []string{"melissa", "sort"}},
			wantErr: true,
		},
		{
			name:    "unHappy Path: sort is not support",
			args:    args{args: []string{"melissa", "--sort-created", "a"}},
			wantErr: true,
		},
		{
			name:    "unHappy Path: username does not exist",
			args:    args{args: []string{"melissa2", "--sort-name", "asc"}},
			wantErr: true,
		},
	}
	InitData()
	CreateUserController([]string{"melissa"})
	CreateFolderController([]string{"melissa", "foldername"})
	CreateFolderController([]string{"melissa", "foldername1"})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ListFolderController(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("ListFolderController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
