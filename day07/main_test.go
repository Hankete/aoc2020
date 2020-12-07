package main

import (
	"reflect"
	"testing"
)

func Test_parseLine1(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		line string
		want Bag
		//want1 string
	}{
		{"1", "dotted black bags contain no other bags.", Bag{}},
		{
			"2",
			"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
			Bag{
				containedBags: []string{"bright white", "muted yellow"},
				counters:      []int64{3, 4},
			},
		},
		{
			"3",
			"bright white bags contain 1 shiny gold bag.",
			Bag{
				containedBags: []string{"shiny gold"},
				counters:      []int64{1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := parseLine(tt.line)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseLine() got = %v, want %v", got, tt.want)
			}
			//if got1 != tt.want1 {
			//	t.Errorf("parseLine() got1 = %v, want %v", got1, tt.want1)
			//}
		})
	}
}
