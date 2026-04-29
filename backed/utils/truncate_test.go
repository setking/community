package utils

import "testing"

func TestTruncateByWords(t *testing.T) {
	type args struct {
		s        string
		maxWords int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "Test case 1", args: args{s: "Hello, world!", maxWords: 2}, want: "Hello, world!"},
		{name: "Test case 2", args: args{s: "This is a test", maxWords: 3}, want: "This is a test"},
		{name: "Test case 3", args: args{s: "This is a very long string", maxWords: 5}, want: "This is a very long string"},
		{name: "Test case 4", args: args{s: "This is a very long string", maxWords: 1}, want: "This"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TruncateByWords(tt.args.s, tt.args.maxWords); got != tt.want {
				t.Errorf("TruncateByWords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSeparator(t *testing.T) {
	type args struct {
		r rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "Test case 1", args: args{r: ','}, want: true},
		{name: "Test case 2", args: args{r: '!'}, want: true},
		{name: "Test case 3", args: args{r: '.'}, want: true},
		{name: "Test case 4", args: args{r: '-'}, want: false},
		{name: "Test case 5", args: args{r: 'a'}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSeparator(tt.args.r); got != tt.want {
				t.Errorf("isSeparator() = %v, want %v", got, tt.want)
			}
		})
	}
}
