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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteFolderController(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("DeleteFolderController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
