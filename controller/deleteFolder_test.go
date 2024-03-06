package controller

import "testing"

func TestDeleteFolderController(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{

		{
			name:    "Happy Path: delete folder successfully.",
			args:    args{args: []string{"melissa", "foldername"}},
			wantErr: false,
		},
		{
			name:    "unHappy Path: delete folder failed: username contains special char",
			args:    args{args: []string{"melissa!", "foldername"}},
			wantErr: true,
		},
		{
			name:    "unHappy Path: delete folder failed: foldername contains special char",
			args:    args{args: []string{"melissa", "foldername!"}},
			wantErr: true,
		},
		{
			name:    "unHappy Path: username does not exist",
			args:    args{args: []string{"melissa2", "foldername"}},
			wantErr: true,
		},
		{
			name:    "unHappy Path: foldername does not exist",
			args:    args{args: []string{"melissa", "foldername1"}},
			wantErr: true,
		},
	}
	InitData()
	CreateUserController([]string{"melissa"})
	CreateFolderController([]string{"melissa", "foldername"})
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteFolderController(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("DeleteFolderController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
