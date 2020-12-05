package main

import "testing"

func Test_findColumn(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		s    string
		want int64
	}{
		{"1", "RRR", 7},
		{"2", "RLL", 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findColumn(tt.s); got != tt.want {
				t.Errorf("findColumn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findRow(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		s    string
		want int64
	}{
		{"1", "BFFFBBF", 70},
		{"2", "FFFBBBF", 14},
		{"3", "BBFFBBF", 102},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findRow(tt.s); got != tt.want {
				t.Errorf("findRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
