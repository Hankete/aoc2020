package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("day03/day03-input.txt")
	if err != nil {
		panic("Can't open file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	board := make([]string, 0)
	for scanner.Scan() {
		massStr := scanner.Text()
		board = append(board, massStr)
	}

	paths := []int{1, 1, 3, 1, 5, 1, 7, 1, 1, 2}
	answers := make([]int, 0)
	for j := 0; j < len(paths); j = j + 2 {
		pathR, pathD := paths[j], paths[j+1]
		x := 0
		sum := 0
		for y := pathD; y < len(board); y += pathD {
			for i := 1; i <= pathR; i++ {
				if x+1 == len(board[y]) {
					x = 0
				} else {
					x += 1
				}
			}
			if board[y][x] == '#' {
				sum += 1
			}
		}
		answers = append(answers, sum)
	}

	answer := 1

	for i := 0; i < len(answers); i++ {
		answer *= answers[i]
	}

	fmt.Println("answer ", answer)

}
