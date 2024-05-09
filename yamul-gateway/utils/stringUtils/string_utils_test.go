package stringUtils

import "testing"

func TestTrimRight(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base case",
			args: args{
				str: "hello world",
			},
			want: "hello world",
		},
		{
			name: "left trim",
			args: args{
				str: "\x00hello world",
			},
			want: "\x00hello world",
		},
		{
			name: "right trim",
			args: args{
				str: "hello world\x00",
			},
			want: "hello world",
		},
		{
			name: "full case",
			args: args{
				str: "\x00\x00\x00hello\x00 world \x00\x00\x00\x00\x00",
			},
			want: "\x00\x00\x00hello\x00 world ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TrimRight(tt.args.str); got != tt.want {
				t.Errorf("TrimRight() = %v, want %v", got, tt.want)
			}
		})
	}
}
