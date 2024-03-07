package main

import (
	"reflect"
	"testing"
)

func Test_processArg(t *testing.T) {
	type args struct {
		input string
		op    string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "Happy Path",
			args: args{
				input: "register melissa",
				op:    "register",
			},
			want:    []string{"melissa"},
			wantErr: false,
		},
		{
			name: "Happy Path",
			args: args{
				input: `register "melissa"`,
				op:    "register",
			},
			want:    []string{`"melissa"`},
			wantErr: false,
		},
		{
			name: "Happy Path",
			args: args{
				input: `register "melissa`,
				op:    "register",
			},
			want:    []string{},
			wantErr: true,
		},
		{
			name: "Happy Path",
			args: args{
				input: `register melissa"`,
				op:    "register",
			},
			want:    []string{},
			wantErr: true,
		},

		{
			name: "Happy Path",
			args: args{
				input: `register "m wang"`,
				op:    "register",
			},
			want:    []string{`"m wang"`},
			wantErr: false,
		},
		{
			name: "Happy Path",
			args: args{
				input: `register "m wang`,
				op:    "register",
			},
			want:    []string{},
			wantErr: true,
		},
		{
			name: "Happy Path",
			args: args{
				input: `register m wang"`,
				op:    "register",
			},
			want:    []string{},
			wantErr: true,
		},

		{
			name: "Happy Path",
			args: args{
				input: `create-folder "m wang" "folder 1"`,
				op:    "create-folder",
			},
			want:    []string{`"m wang"`, `"folder 1"`},
			wantErr: false,
		},

		{
			name: "Happy Path",
			args: args{
				input: `create-folder "m wang" --sort-name asc`,
				op:    "create-folder",
			},
			want:    []string{`"m wang"`, `--sort-name`, `asc`},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := processArg(tt.args.input, tt.args.op)
			if (err != nil) != tt.wantErr {
				t.Errorf("processArg() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("processArg() got = %v, want %v", got, tt.want)
			}
		})
	}
}
