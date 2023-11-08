package server

import (
	"testing"
)

func Test_shortHash(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{"https://github.com/avelino/awesome-go"},
			want: "7506107d43",
		},
		{
			name: "test2",
			args: args{"erftyhujikolijuhygtfveyurijvoer"},
			want: "4813e06f2e",
		},
		{
			name: "test3",
			args: args{"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat."},
			want: "1c81c608a6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortHash(tt.args.s); got != tt.want {
				t.Errorf("shortHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShortURLHashed(t *testing.T) {
	type args struct {
		longURL string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{"https://github.com/avelino/awesome-go"},
			want: "https://urlshorter/7506107d43",
		},
		{
			name: "test2",
			args: args{"erftyhujikolijuhygtfveyurijvoer"},
			want: "https://urlshorter/4813e06f2e",
		},
		{
			name: "test3",
			args: args{"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat."},
			want: "https://urlshorter/1c81c608a6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ShortURLHashed(tt.args.longURL); got != tt.want {
				t.Errorf("ShortURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
