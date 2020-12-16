package main

import (
	"AOC2020/ioutils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func solve1() {
	path := "day15/test.txt"
	lines := ioutils.ReadLines(path)
	parts := strings.Split(lines[0], ",")
	startingNumbers := make([]int64, 0)
	usage := make(map[int64]int64)
	lSpoke := int64(0)
	for _, part := range parts {
		n, _ := strconv.ParseInt(part, 10, 64)
		startingNumbers = append(startingNumbers, n)
	}
	var n int64
	for i := int64(1); i <= 30000000; i++ {

		if int(i)-1 < len(startingNumbers) {
			n = startingNumbers[i-1]
		} else {
			if _, ok := usage[lSpoke]; !ok {
				n = 0
			} else {
				n = i - 1 - usage[lSpoke]
			}
		}
		if i > 1 {
			usage[lSpoke] = i - 1

		}
		lSpoke = n
		//fmt.Println(n)
	}
	fmt.Println(n)

}

func solve2() {
	//path := "day14/input.txt"
	//lines := ioutils.ReadLines(path)

}

func main() {
	start := time.Now()
	solve1()
	fmt.Println(time.Now().Sub(start))

}
