package main

import (
	"AOC2020/ioutils"
	"fmt"
	"strconv"
)

func solve1() {
	path := "day12/day12-input.txt"
	lines := ioutils.ReadLines(path)

	x1, y1 := 0, 0
	x2, y2 := 10, 1
	directions := make(map[string][]int, 0)
	directions["W"] = []int{-1, 0}
	directions["E"] = []int{1, 0}
	directions["N"] = []int{0, 1}
	directions["S"] = []int{0, -1}

	currDir := "E"

	for _, line := range lines {
		direction := line[:1]
		val64, _ := strconv.ParseInt(line[1:], 10, 64)
		val := int(val64)

		x1, y1, x2, y2, currDir = move(x1, y1, x2, y2, currDir, direction, val, directions)
		fmt.Println(x1, y1, abs(x1, y1))
	}
}

func abs(x, y int) int {
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	return x + y
}

func move(x1, y1, x2, y2 int, currDirection, direction string, val int, directions map[string][]int) (int, int, int, int, string) {
	if direction == "E" || direction == "W" || direction == "N" || direction == "S" {
		//currDirection = direction
		x2 += directions[direction][0] * val
		y2 += directions[direction][1] * val
		return x1, y1, x2, y2, currDirection
	} else if direction == "F" {
		x1 += val * x2
		y1 += val * y2
		return x1, y1, x2, y2, currDirection
	} else if direction == "L" || direction == "R" {

		//currDirection = newDirection
		x2, y2 = direct(x2, y2, val, direction)
		//x += directions[newDirection][0] * val
		//y += directions[newDirection][1] * val
		return x1, y1, x2, y2, currDirection
	}
	panic("Wrong")
}

func direct(x, y, val int, turn string) (int, int) {

	times := val / 90

	for times > 0 {
		if turn == "L" {
			x, y = -y, x
		} else if turn == "R" {
			x, y = y, -x
		}
		times--
	}

	return x, y
}

func main() {
	solve1()
}
