package utils

import "testing"

func TestUint64ToInt64Safe(t *testing.T) {
	type args struct {
		u uint64
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
		{name: "Test Case 1", args: args{u: 12345678}, want: 12345678, wantErr: false},
		{name: "Test Case 2", args: args{u: 0}, want: 0, wantErr: false},
		{name: "Test Case 3", args: args{u: 1844674407}, want: 1844674407, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Uint64ToInt64Safe(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("Uint64ToInt64Safe() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Uint64ToInt64Safe() = %v, want %v", got, tt.want)
			}
		})
	}
}
