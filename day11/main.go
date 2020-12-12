package main

import (
	"AOC2020/ioutils"
	"fmt"
)

func solve1() {
	path := "day11/day11-input.txt"
	lines := ioutils.ReadLines(path)
	changed := true
	board := lines
	for changed == true {
		board, changed = occupy(board)
	}
	counter := 0
	for _, line := range board {
		for i := 0; i < len(line); i++ {
			if line[i] == '#' {
				counter += 1
			}
		}
	}
	fmt.Println(board, counter)
}

func occupy(board []string) ([]string, bool) {
	newBoard := make([]string, len(board))
	changed := false
	countOcc := 0
	for i := 0; i < len(board); i++ {
		newRow := ""
		for j := 0; j < len(board[i]); j++ {
			occ := checkAdj(board, i, j)
			if board[i][j] == '.' {
				newRow += "."
			} else if board[i][j] == 'L' && occ == 0 {
				newRow += "#"
				countOcc += 1
				changed = true
			} else if board[i][j] == '#' && occ >= 5 {
				newRow += "L"
				changed = true
			} else {
				newRow += string(board[i][j])
			}
			newBoard[i] = newRow
		}
	}
	return newBoard, changed
}

func IsOccupied(board []string, i, j, v1, v2 int) int {
	i += v1
	j += v2
	for checkIndexes(board, i, j) {
		if board[i][j] == '#' {
			return 1
		}
		if board[i][j] == 'L' {
			return 0
		}
		i += v1
		j += v2
	}
	return 0
}

func checkAdj(board []string, i, j int) int {
	counterOcc := 0
	counterOcc += IsOccupied(board, i, j, -1, -1)
	counterOcc += IsOccupied(board, i, j, -1, 0)
	counterOcc += IsOccupied(board, i, j, 0, -1)
	counterOcc += IsOccupied(board, i, j, 1, 0)
	counterOcc += IsOccupied(board, i, j, 0, 1)
	counterOcc += IsOccupied(board, i, j, 1, 1)
	counterOcc += IsOccupied(board, i, j, 1, -1)
	counterOcc += IsOccupied(board, i, j, -1, 1)

	return counterOcc
}

func checkIndexes(board []string, i, j int) bool {
	return i >= 0 && j >= 0 && i < len(board) && j < len(board[i])
}
func main() {
	solve1()
}
