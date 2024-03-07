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
			},
			want:    []string{`register`, `melissa`},
			wantErr: false,
		},
		{
			name: "Happy Path",
			args: args{
				input: `register "melissa"`,
			},
			want:    []string{`register`, `"melissa"`},
			wantErr: false,
		},
		{
			name: "unHappy Path:missing a \"",
			args: args{
				input: `register "melissa`,
			},
			want:    []string{},
			wantErr: true,
		},
		{
			name: "unHappy Path:missing a \"",
			args: args{
				input: `register melissa"`,
			},
			want:    []string{},
			wantErr: true,
		},
		{
			name: "Happy Path",
			args: args{
				input: `register "m wang"`,
			},
			want:    []string{`register`, `"m wang"`},
			wantErr: false,
		},
		{
			name: "unHappy Path:missing a \"",
			args: args{
				input: `register "m wang`,
			},
			want:    []string{},
			wantErr: true,
		},
		{
			name: "unHappy Path:missing a \"",
			args: args{
				input: `register m wang"`,
			},
			want:    []string{},
			wantErr: true,
		},

		{
			name: "Happy Path",
			args: args{
				input: `create-folder "m wang" "folder 1"`,
			},
			want:    []string{`create-folder`, `"m wang"`, `"folder 1"`},
			wantErr: false,
		},

		{
			name: "Happy Path",
			args: args{
				input: `create-folder "m wang" --sort-name asc`,
			},
			want:    []string{`create-folder`, `"m wang"`, `--sort-name`, `asc`},
			wantErr: false,
		},
		{
			name: "Happy Path",
			args: args{
				input: `create-file "m wang" "filename 1" --sort-name asc`,
			},
			want:    []string{`create-file`, `"m wang"`, `"filename 1"`, `--sort-name`, `asc`},
			wantErr: false,
		},
		{
			name: "Happy Path",
			args: args{
				input: `create-file "m wang" filename1 --sort-name asc`,
			},
			want:    []string{`create-file`, `"m wang"`, `filename1`, `--sort-name`, `asc`},
			wantErr: false,
		},

		{
			name: "Happy Path",
			args: args{
				input: `list-file "m wang" "foldername1" "filename 1" --sort-name asc`,
			},
			want:    []string{`list-file`, `"m wang"`, `"foldername1"`, `"filename 1"`, `--sort-name`, `asc`},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := processArg(tt.args.input)
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
