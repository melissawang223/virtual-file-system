package helper

import (
	"reflect"
	"testing"
)

func TestCheckFileName(t *testing.T) {
	type args struct {
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Input looks good",
			args:    args{"fileName"},
			wantErr: false,
		},
		{
			name:    "Unhappy Path: Input too long",
			args:    args{"fileNamefileNamefileNamefileNamefileNamefileName"},
			wantErr: true,
		},
		{
			name:    "Unhappy Path: Input is invalid",
			args:    args{"fileName!!"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckFileName(tt.args.fileName); (err != nil) != tt.wantErr {
				t.Errorf("CheckFileName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckFileDescription(t *testing.T) {
	type args struct {
		fileDes string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Happy Path: description  looks good",
			args:    args{"description example"},
			wantErr: false,
		},
		{
			name:    "Unhappy Path: Input too long",
			args:    args{"fileNamefileNamefileNamefileNamefileNamefileName"},
			wantErr: true,
		},
		{
			name:    "Happy Path: description is empty",
			args:    args{""},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckFileDescription(tt.args.fileDes); (err != nil) != tt.wantErr {
				t.Errorf("CheckFileDescription() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckFolderName(t *testing.T) {
	type args struct {
		folderName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Input looks good",
			args:    args{"folderName"},
			wantErr: false,
		},
		{
			name:    "Unhappy Path: Input too long",
			args:    args{"fileNamefileNamefileNamefileNamefileNamefolderName"},
			wantErr: true,
		},
		{
			name:    "Unhappy Path: Input is invalid",
			args:    args{"folderName!!"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckFolderName(tt.args.folderName); (err != nil) != tt.wantErr {
				t.Errorf("CheckFolderName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckUserName(t *testing.T) {
	type args struct {
		userName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Happy Path: Input looks good",
			args:    args{"melissa"},
			wantErr: false,
		},
		{
			name:    "Unhappy Path: Input too long",
			args:    args{"melissa1234"},
			wantErr: true,
		},
		{
			name:    "Unhappy Path: Input is invalid",
			args:    args{"melissa!"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckUserName(tt.args.userName); (err != nil) != tt.wantErr {
				t.Errorf("CheckUserName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_CheckArguments(t *testing.T) {
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
			got, err := CheckArguments(tt.args.input)
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
