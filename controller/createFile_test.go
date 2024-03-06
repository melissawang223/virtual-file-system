package controller

import "testing"

func TestCreateFileController(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Happy Path: create file successfully.",
			args:    args{args: []string{"melissa", "foldername", "filename", "description"}},
			wantErr: false,
		},
		{
			name:    "Happy Path: create file without description successfully.",
			args:    args{args: []string{"melissa", "foldername", "filename1"}},
			wantErr: false,
		},
		{
			name:    "unHappy Path: create file failed: foldername contains special char",
			args:    args{args: []string{"melissa", "foldername!", "filename3", "description"}},
			wantErr: true,
		},
		{
			name:    "unHappy Path: username does not exist",
			args:    args{args: []string{"melissa2", "foldername", "filename4", "description"}},
			wantErr: true,
		},
		{
			name:    "unHappy Path: create file failed: filename contains special char",
			args:    args{args: []string{"melissa", "foldername", "filename5!", "description"}},
			wantErr: true,
		},
		{
			name:    "unHappy Path: foldername does not exist",
			args:    args{args: []string{"melissa2", "notexsitfoldername", "filename", "description"}},
			wantErr: true,
		},
	}

	CreateUserController([]string{"melissa"})
	CreateFolderController([]string{"melissa", "foldername"})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateFileController(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("CreateFileController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
