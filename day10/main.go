package main

import (
	"AOC2020/ioutils"
	"fmt"
	"sort"
	"strconv"
)

func findMax(numbers []int64) int64 {
	max := int64(0)
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	return max
}

func solve1() {
	path := "day10/day10-input.txt"
	lines := ioutils.ReadLines(path)

	jolts := make([]int64, 0)

	for _, line := range lines {
		jolt, _ := strconv.ParseInt(line, 10, 64)
		jolts = append(jolts, jolt)
	}

	sort.Slice(jolts, func(i, j int) bool { return jolts[i] < jolts[j] })

	lastJolt := int64(0)
	counters := make(map[int64]int64)
	for _, jolt := range jolts {
		counters[jolt-lastJolt]++
		lastJolt = jolt
	}

	counters[3]++
	answer := counters[3] * counters[1]
	fmt.Println(answer)
}

func isValid(numbers []int64, myJolt int64) bool {
	if len(numbers) == 0 || numbers[0] > 3 || myJolt-numbers[len(numbers)-1] <= 3 {
		return false
	}
	for i := 1; i < len(numbers); i++ {
		if numbers[i]-numbers[i-1] > 3 {
			return false
		}
	}
	return true
}

func countRecursive(series, numbers []int64, j int, max int64) int {
	if series[len(series)-1] == max {
		//fmt.Println(series)
		return 1
	}
	ls := series[len(series)-1]
	answer := 0
	var newSeries []int64
	if j+1 < len(numbers) && numbers[j+1]-ls <= 3 {
		newSeries = append(series, numbers[j+1])
		answer += countRecursive(newSeries, numbers, j+1, max)
	}
	if j+2 < len(numbers) && numbers[j+2]-ls <= 3 {
		newSeries = append(series, numbers[j+2])
		answer += countRecursive(newSeries, numbers, j+2, max)
	}
	if j+3 < len(numbers) && numbers[j+3]-ls <= 3 {
		newSeries = append(series, numbers[j+3])
		answer += countRecursive(newSeries, numbers, j+3, max)
	}

	return answer

}

func solve2() {
	path := "day10/day10-input.txt"
	lines := ioutils.ReadLines(path)

	jolts := make([]int64, 0)

	for _, line := range lines {
		jolt, _ := strconv.ParseInt(line, 10, 64)
		jolts = append(jolts, jolt)
	}
	jolts = append(jolts, 0)
	sort.Slice(jolts, func(i, j int) bool { return jolts[i] < jolts[j] })
	max := jolts[len(jolts)-1] + 3
	jolts = append(jolts, max)

	fmt.Println(jolts)
	series := make([]int64, 0)
	series = append(series, 0)
	canBeMoved := make(map[int64]int64)
	canBeMoved[0] = 1
	for _, i := range jolts[1:] {
		canBeMoved[i] = canBeMoved[i-1] + canBeMoved[i-2] + canBeMoved[i-3]
	}

	fmt.Println(canBeMoved)
}

func main() {
	solve2()
}
