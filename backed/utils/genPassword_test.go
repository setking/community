package utils

import "testing"

func TestGenMd5(t *testing.T) {
	type args struct {
		pwd string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{pwd: "password"}, want: "$pbkdf2-sha512$WF38GynfLi9T88Y2$1431b7deee8594d9418ef70faf63e59aae2e4bb0e34452a37069d94afc4aa9b5"},
		{name: "test2", args: args{pwd: "123456"}, want: "$pbkdf2-sha512$fEt75S87Nj1anAKT$791f328a947ceebd2ca6a9901161864bd80f9fee2a5e215214ccbe9e28b7136c"},
		{name: "test3", args: args{pwd: "abc123"}, want: "$pbkdf2-sha512$pN2ivtOrwWEdwkXX$aad91a07b04dd692d4157af3e97e240f2cd8dd57947db3878a26f94ba9b9199d"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenMd5(tt.args.pwd); got != tt.want {
				t.Errorf("GenMd5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerifyPassword(t *testing.T) {
	type args struct {
		genPwd string
		oldPwd string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{name: "test1", args: args{genPwd: "$pbkdf2-sha512$WF38GynfLi9T88Y2$1431b7deee8594d9418ef70faf63e59aae2e4bb0e34452a37069d94afc4aa9b5", oldPwd: "password"}, want: true},
		{name: "test2", args: args{genPwd: "$pbkdf2-sha512$fEt75S87Nj1anAKT$791f328a947ceebd2ca6a9901161864bd80f9fee2a5e215214ccbe9e28b7136c", oldPwd: "123456"}, want: true},
		{name: "test3", args: args{genPwd: "$pbkdf2-sha512$pN2ivtOrwWEdwkXX$aad91a07b04dd692d4157af3e97e240f2cd8dd57947db3878a26f94ba9b9199d", oldPwd: "abc123"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VerifyPassword(tt.args.genPwd, tt.args.oldPwd); got != tt.want {
				t.Errorf("VerifyPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
