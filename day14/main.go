package main

import (
	"AOC2020/ioutils"
	"fmt"
	"strconv"
	"strings"
)

func solve1() {
	path := "day14/day14-input.txt"
	lines := ioutils.ReadLines(path)
	mask := ""
	sum := int64(0)
	m := make(map[string]int64, 0)
	for _, line := range lines {
		parts := strings.Split(line, " ")
		if parts[0] == "mask" {
			mask = parts[2]
		}
		var l1, l2 int
		var newString string
		if strings.HasPrefix(parts[0], "mem") {
			for i := 1; len(mask)-i >= 0; i++ {

				var newChar string
				n, _ := strconv.ParseInt(parts[2], 10, 64)
				bits := strconv.FormatInt(n, 2) // 1111011
				l1, l2 = len(bits)-i, len(mask)-i
				if l1 < 0 {
					newChar = "0"
				} else {
					newChar = string(bits[l1])
				}

				if mask[l2] == 'X' {
					newString = newChar + newString
				} else {
					newString = string(mask[l2]) + newString
				}

			}
		}
		key := parts[0][3:]
		parsed, _ := strconv.ParseInt(newString, 2, 64)
		fmt.Println(newString, parsed)
		m[key] = parsed

	}

	for _, val := range m {
		sum += val
		fmt.Println(sum)
	}

}

func solve2() {
	path := "day14/day14-input.txt"
	lines := ioutils.ReadLines(path)
	mask := ""
	sum := int64(0)
	m := make(map[string]int64, 0)
	for _, line := range lines {
		parts := strings.Split(line, " ")
		if parts[0] == "mask" {
			mask = parts[2]
		}

		var l1, l2 int
		var newString string
		if strings.HasPrefix(parts[0], "mem") {
			key := strings.TrimPrefix(strings.TrimSuffix(parts[0], "]"), "mem[")
			key2, _ := strconv.ParseInt(key, 10, 64)
			keyToBin := strconv.FormatInt(key2, 2)
			fmt.Println(key2, keyToBin)
			for i := 1; len(mask)-i >= 0; i++ {
				var newChar string

				l1, l2 = len(keyToBin)-i, len(mask)-i
				if l1 < 0 {
					newChar = "0"
				} else {
					newChar = string(keyToBin[l1])
				}

				if mask[l2] == 'X' {
					newString = "X" + newString
				} else if mask[l2] == '0' {
					newString = newChar + newString
				} else if mask[l2] == '1' {
					newString = "1" + newString
				}

			}
			parsed, _ := strconv.ParseInt(parts[2], 10, 64)
			fmt.Println(newString, parsed)
			SavetoMap(m, newString, parsed)
		}

	}

	for _, val := range m {
		sum += val
		fmt.Println(sum)
	}
}

func SavetoMap(m map[string]int64, key string, val int64) {
	idx := strings.IndexRune(key, 'X')
	if idx == -1 {
		m[key] = val
	}
	if idx != -1 {
		newKey1 := key[:idx] + "0" + key[idx+1:]
		newKey2 := key[:idx] + "1" + key[idx+1:]

		SavetoMap(m, newKey1, val)
		SavetoMap(m, newKey2, val)

	}
}

func main() {
	solve2()
}
