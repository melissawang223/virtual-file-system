package controller

import (
	"testing"
)

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
	InitData()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateUserController(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("CreateUserController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
