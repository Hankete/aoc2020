package main

import (
	"AOC2020/ioutils"
	"fmt"
	"strconv"
)

func doOperations(lines []string, idx int64) bool {
	i := int64(0)
	visited := make(map[int64]bool)
	accumulator := int64(0)
	for {
		if int(i) == len(lines) {
			fmt.Println(accumulator)
			return true
		}
		if visited[i] == true {
			return false
		}
		line := lines[i]
		op, valStr := line[:3], line[4:]
		visited[i] = true
		if i == idx {
			if op == "nop" {
				op = "jmp"
			} else if op == "jmp" {
				op = "nop"
			}
		}
		if op == "nop" {
			i++
			continue
		}
		if op == "acc" {
			val, _ := strconv.ParseInt(valStr[1:], 10, 64)
			if valStr[0] == '+' {
				i++
				accumulator += val
			} else {
				i++
				accumulator -= val
			}
		}
		if op == "jmp" {
			val, _ := strconv.ParseInt(valStr[1:], 10, 64)
			if valStr[0] == '+' {
				i += val
			} else {
				i -= val
			}
		}

	}

	return true
}

func main() {
	path := "day08/day08-input.txt"
	lines := ioutils.ReadLines(path)

	for idx, _ := range lines {
		linesCopy := ioutils.ReadLines(path)
		if doOperations(linesCopy, int64(idx)) {
			fmt.Println("FOUND")
		}
	}

}
