package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isValid(str string, min, max int, letter rune) bool {
	counter := 0
	for _, val := range str {
		if val == letter {
			counter += 1
		}
	}

	return counter >= min && counter <= max
}

func isValid2(str string, min, max int, letter rune) bool {
	counter := 0
	length := len(str)
	if min-1 < length && str[min-1] == uint8(letter) {
		counter += 1
	}
	if max-1 < length && str[max-1] == uint8(letter) {
		counter += 1
	}

	if counter == 1 {
		return true
	}

	return false
}

func main() {
	f, err := os.Open("day02/day02-input.txt")
	if err != nil {
		panic("Can't open file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	sum := 0
	for scanner.Scan() {
		massStr := scanner.Text()
		parts := strings.Split(massStr, " ")
		ranges := strings.Split(parts[0], "-")
		min, _ := strconv.ParseInt(ranges[0], 10, 0)
		max, _ := strconv.ParseInt(ranges[1], 10, 0)
		letter := parts[1][0]
		word := parts[2]
		fmt.Println(min, max, letter, word)
		if isValid2(word, int(min), int(max), rune(letter)) {
			sum += 1
		}
	}

	fmt.Println("andswer ", sum)

}
