package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calcFuel(mass int64) int64 {
	sum := int64(0)
	for fuel := mass/3 - 2; fuel > 0; {
		sum += fuel
		fuel = fuel/3 - 2
	}
	return sum
}

func main() {
	f, err := os.Open("day01/day01-input.txt")
	if err != nil {
		panic("Can't open file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	sum := 0
	fuel := 0
	list := make([]int64, 0)
	for scanner.Scan() {
		massStr := scanner.Text()
		mass, _ := strconv.ParseInt(massStr, 10, 0)
		list = append(list, mass)
		sum += fuel
	}

	for i := 0; i < len(list); i++ {
		for j := i; j < len(list); j++ {
			for k := j; k < len(list); k++ {
				if list[i]+list[j]+list[k] == 2020 {
					fmt.Println("Numbers", list[i], list[j], list[k], " answer ", list[i]*list[j]*list[k])
				}
			}

		}
	}
}
