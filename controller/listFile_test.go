package controller

import "testing"

func TestListFileController(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Happy Path: list-file --sort-name asc successfully.",
			args:    args{args: []string{"melissa", "foldername", "--sort-name", "asc"}},
			wantErr: false,
		},
		{
			name:    "Happy Path: list-file --sort-created asc successfully.",
			args:    args{args: []string{"melissa", "foldername", "--sort-created", "asc"}},
			wantErr: false,
		},
		{
			name:    "unHappy Path: sort is not support",
			args:    args{args: []string{"melissa", "foldername", "sort"}},
			wantErr: true,
		},
		{
			name:    "unHappy Path: sort is not support",
			args:    args{args: []string{"melissa", "foldername", "--sort-created", "a"}},
			wantErr: true,
		},
		{
			name:    "unHappy Path: username does not exist",
			args:    args{args: []string{"melissa2", "foldername", "--sort-name", "asc"}},
			wantErr: true,
		},
		{
			name:    "unHappy Path: foldername does not exist",
			args:    args{args: []string{"melissa", "notexitfoldername", "--sort-name", "asc"}},
			wantErr: true,
		},
	}

	CreateUserController([]string{"melissa"})
	CreateFolderController([]string{"melissa", "foldername"})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ListFileController(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("ListFileController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
