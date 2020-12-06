package ioutils

import (
	"bufio"
	"os"
)

func ReadLines(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		panic("Can't open file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := make([]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func ReadLinesByBlank(path string) [][]string {
	lines := ReadLines(path)
	groups := make([][]string, 0)

	i := 0
	for idx, line := range lines {
		if line == "" && idx != len(lines) {
			groups = append(groups, lines[i:idx])
			i = idx + 1
		}
	}

	return groups
}
