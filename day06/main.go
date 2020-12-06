package main

import (
	"AOC2020/ioutils"
	"fmt"
)

func countGroup(group []string) int {
	yes := make(map[string]bool)
	for _, answers := range group {
		for _, answer := range answers {
			yes[string(answer)] = true
		}
	}
	counter := 0
	for _, _ = range yes {
		counter += 1
	}
	return counter
}

func countGroup2(group []string) int {
	yes := make(map[string]int)
	for _, answers := range group {
		for _, answer := range answers {
			yes[string(answer)] += 1
		}
	}
	counter := 0
	for _, val := range yes {
		if len(group) == val {
			counter += 1
		}
	}
	return counter
}

func main() {
	groups := ioutils.ReadLinesByBlank("day06/day06-input.txt")

	counter := 0

	for _, group := range groups {
		counter += countGroup2(group)
	}
	fmt.Println(counter)
}
