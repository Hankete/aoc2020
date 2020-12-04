package main

import "testing"

func Test_isValid(t *testing.T) {
	type args struct {
		str    string
		min    int
		max    int
		letter rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"abcde", 1, 3, 'a'}, true},
		{"2", args{"cdefg", 1, 3, 'b'}, false},
		{"3", args{"ccccccccc", 2, 9, 'c'}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.str, tt.args.min, tt.args.max, tt.args.letter); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValid2(t *testing.T) {
	type args struct {
		str    string
		min    int
		max    int
		letter rune
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1", args{"abcde", 1, 3, 'a'}, true},
		{"2", args{"cdefg", 1, 3, 'b'}, false},
		{"3", args{"ccccccccc", 2, 9, 'c'}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid2(tt.args.str, tt.args.min, tt.args.max, tt.args.letter); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
