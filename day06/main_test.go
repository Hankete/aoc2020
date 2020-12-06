package main

import "testing"

func Test_countGroup(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name  string
		group []string
		want  int
	}{
		{"1", []string{"abcx", "abcy", "abcz"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countGroup(tt.group); got != tt.want {
				t.Errorf("countGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
