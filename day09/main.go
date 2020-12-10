package main

import (
	"AOC2020/ioutils"
	"fmt"
	"strconv"
)

func isPossible(number int64, possible []int64) bool {
	for _, poss := range possible {
		if poss == number {
			return true
		}
	}
	return false
}

func makePossibles(prevs []int64) []int64 {
	possible := make([]int64, 0)

	for i := 0; i < len(prevs); i++ {
		for j := i + 1; j < len(prevs); j++ {
			possible = append(possible, prevs[i]+prevs[j])
		}
	}
	return possible
}

func solve1() {
	path := "day09/day09-input.txt"
	preambleSize := 5
	lines := ioutils.ReadLines(path)
	numbers := make([]int64, 0)
	for _, line := range lines {
		number, _ := strconv.ParseInt(line, 10, 64)
		numbers = append(numbers, number)
	}
	prevs := numbers[:preambleSize]
	numbers = numbers[preambleSize:]

	for _, number := range numbers {
		possible := makePossibles(prevs)
		if !isPossible(number, possible) {
			fmt.Println(number)
			break
		}
		prevs = prevs[1:]
		prevs = append(prevs, number)
	}
}

func MinMax(array []int64) (int64, int64) {
	var max int64 = array[0]
	var min int64 = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}
func solve2() {
	path := "day09/day09-input.txt"
	lines := ioutils.ReadLines(path)
	numbers := make([]int64, 0)
	searching := int64(104054607)
	for _, line := range lines {
		number, _ := strconv.ParseInt(line, 10, 64)
		numbers = append(numbers, number)
	}

	i, j := 0, 0
	actual := numbers[i]
	for j < len(numbers) {
		if actual < searching {
			j++
			actual += numbers[j]
		} else if actual > searching {
			actual -= numbers[i]

			i++
		} else {
			fmt.Println(numbers[i:j+1], i, j)
			fmt.Println(MinMax(numbers[i : j+1]))
			break
		}

	}
}

// 104054607
func main() {
	solve2()
}
