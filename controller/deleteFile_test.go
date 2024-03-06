package controller

import "testing"

func TestDeleteFileController(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Happy Path: delete file successfully.",
			args:    args{args: []string{"melissa", "foldername", "filename"}},
			wantErr: false,
		},
		{
			name:    "unHappy Path: delete file failed: username contains special char",
			args:    args{args: []string{"melissa!", "foldername", "filename"}},
			wantErr: true,
		},
		{
			name:    "unHappy Path: delete file failed: foldername contains special char",
			args:    args{args: []string{"melissa", "foldername!", "filename"}},
			wantErr: true,
		},
		{
			name:    "unHappy Path: delete file failed: filename contains special char",
			args:    args{args: []string{"melissa", "foldername", "filename!"}},
			wantErr: true,
		},
		{
			name:    "unHappy Path: username does not exist",
			args:    args{args: []string{"melissa2", "foldername", "filename"}},
			wantErr: true,
		},
		{
			name:    "unHappy Path: foldername does not exist",
			args:    args{args: []string{"melissa", "foldername1", "filename"}},
			wantErr: true,
		},
		{
			name:    "unHappy Path: filename does not exist",
			args:    args{args: []string{"melissa", "foldername", "filename1"}},
			wantErr: true,
		},
	}
	InitData()
	CreateUserController([]string{"melissa"})
	CreateFolderController([]string{"melissa", "foldername"})
	CreateFileController([]string{"melissa", "foldername", "filename"})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteFileController(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("DeleteFileController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
