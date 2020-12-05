package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func findRow(s string) int64 {
	binStr := ""
	for _, val := range s {
		if val == 'F' {
			binStr += "0"
		}
		if val == 'B' {
			binStr += "1"
		}
	}
	answer, _ := strconv.ParseInt(binStr, 2, 64)
	return answer
}

func findColumn(s string) int64 {
	binStr := ""
	for _, val := range s {
		if val == 'L' {
			binStr += "0"
		}
		if val == 'R' {
			binStr += "1"
		}
	}
	answer, _ := strconv.ParseInt(binStr, 2, 64)
	return answer
}

func main() {
	f, err := os.Open("day05/day05-input.txt")
	if err != nil {
		panic("Can't open file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	ids := make([]int64, 0)
	lines := make([]string, 0)
	for scanner.Scan() {
		massStr := scanner.Text()
		lines = append(lines, massStr)
		rowStr := massStr[:7]
		rowID := findRow(rowStr)
		colStr := massStr[7:]
		columnID := findColumn(colStr)

		id := rowID*8 + columnID
		ids = append(ids, id)
	}

	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })
	for i := 1; i < len(ids); i++ {
		if ids[i] == ids[i-1]+2 {
			fmt.Println(ids[i] - 1)
		}
	}
}
