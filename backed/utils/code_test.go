package utils

import (
	"testing"
)

func TestMyCode_Msg(t *testing.T) {
	tests := []struct {
		name string
		c    MyCode
		want string
	}{
		// TODO: Add test cases.
		{name: "Default", c: MyCode(1000), want: "success"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Msg(); got != tt.want {
				t.Errorf("MyCode.Msg() = %v, want %v", got, tt.want)
			}
		})
	}
}
