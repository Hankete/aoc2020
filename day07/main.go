package main

import (
	"AOC2020/ioutils"
	"fmt"
	"strconv"
	"strings"
)

type Bag struct {
	containedBags []string
	counters      []int64
	howMany       int
}

func parseLine(line string) (Bag, string) {
	splittedLine := strings.Split(line, "contain")
	key := strings.TrimSuffix(splittedLine[0], " bags ")

	container := splittedLine[1:]
	if len(container) != 1 {
		panic("Container length > 0")
	}
	bags := strings.Split(strings.Join(container, ""), ",")
	bag := Bag{}
	for _, str := range bags {
		splitedBag := strings.Split(str, " ")
		if str[1] < '0' || str[1] > '9' {
			fmt.Println(str)
			return Bag{}, key
		} else {
			bagID := splitedBag[2] + " " + splitedBag[3]
			counter, _ := strconv.ParseInt(splitedBag[1], 10, 64)
			bag.containedBags = append(bag.containedBags, bagID)
			bag.counters = append(bag.counters, counter)
		}
	}
	return bag, key
}

func countShiny(bags map[string]Bag, visited map[string]bool, bagID string) int {
	if bagID == "shiny gold" {
		return 1
	}
	for _, val := range bags[bagID].containedBags {
		if countShiny(bags, visited, val) > 0 {
			return 1
		}
	}
	return 0
}

func countShiny2(bags map[string]Bag, visited map[string]bool, bagID string) int64 {
	if len(bags[bagID].containedBags) == 0 {
		return 1
	}
	answer := int64(0)
	for idx, val := range bags[bagID].containedBags {
		answer += bags[bagID].counters[idx] * countShiny2(bags, visited, val)
	}
	return answer + 1
}

func main() {
	lines := ioutils.ReadLines("day07/day07-input.txt")
	bags := make(map[string]Bag)
	for _, line := range lines {
		bag, key := parseLine(line)
		bags[key] = bag
	}
	answer := 0
	visited := make(map[string]bool)
	for key, _ := range bags {
		if key != "shiny gold" {
			answer += countShiny(bags, visited, key)
		}
	}

	fmt.Println(answer)
	answer2 := countShiny2(bags, visited, "shiny gold") - 1
	fmt.Println(answer2)
}
