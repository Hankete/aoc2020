package main

import (
	"AOC2020/ioutils"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type rang struct {
	min1 int64
	max1 int64
	min2 int64
	max2 int64
}

func solve1() {
	path := "day16/input.txt"
	lines := ioutils.ReadLinesByBlank(path)
	m := make([]rang, 0)
	for _, part := range lines[0] {
		ranges := strings.Split(part, ":")
		rag := splitRanges(strings.Split(ranges[1], " "))
		fmt.Println(rag)
		m = append(m, rag)
	}
	sum := int64(0)
	for idx, part := range lines[2] {
		if idx == 0 {
			continue
		}

		parts := strings.Split(part, ",")
		for _, part := range parts {
			a, _ := strconv.ParseInt(part, 10, 64)
			if !validNumber(a, m) {
				sum += a
			}
		}

	}
	fmt.Println(sum)
}

func solve2() {
	path := "day16/input.txt"
	lines := ioutils.ReadLinesByBlank(path)
	m := make([]rang, 0)
	for _, part := range lines[0] {
		ranges := strings.Split(part, ":")
		rag := splitRanges(strings.Split(ranges[1], " "))
		fmt.Println(rag)
		m = append(m, rag)
	}
	validNumbers := make([][]int64, 0)
	k := 0
	for idx, part := range lines[2] {
		if idx == 0 {
			continue
		}
		parts := strings.Split(part, ",")
		valid := true
		ticket := make([]int64, 0)
		for _, part := range parts {
			a, _ := strconv.ParseInt(part, 10, 64)
			ticket = append(ticket, a)
			if !validNumber(a, m) {
				valid = false
			}
		}
		if valid {
			validNumbers = append(validNumbers, ticket)
			k += 1
		}

	}
	good := make([][]int, len(validNumbers[0]))

	for i := 0; i < len(validNumbers[0]); i++ {
		good[i] = make([]int, 0)
		for idx, rag := range m {
			valid := true
			for j := 0; j < k; j++ {
				a := validNumbers[j][i]
				if (a >= rag.min1 && a <= rag.max1) || (a >= rag.min2 && a <= rag.max2) {
				} else {
					valid = false
				}
			}
			if valid {
				good[i] = append(good[i], idx)
			}
		}
	}
	//sort.Slice(good, func(i, j int) bool { return len(good[i]) < len(good[j]) })
	fmt.Println(good)
	g := make([]map[int]bool, len(good))
	for i, val := range good {
		g[i] = make(map[int]bool)
		for _, val := range val {
			g[i][val] = true
		}
		fmt.Println(i, len(good[i]), good[i])
	}
	mt := make([]int64, 0)
	numbers := strings.Split(lines[1][1], ",")
	for _, part := range numbers {
		a, _ := strconv.ParseInt(part, 10, 64)
		mt = append(mt, a)
	}
	sum := int64(1)
	// Permutation calculated by hand
	indices2 := []int{15, 2, 6, 13, 14, 1}
	for _, idx := range indices2 {
		sum *= mt[idx]
	}
	fmt.Println(sum)

}

//func generatePerm(g []map[int]bool) []int {
//	answer := make([]int, len(g))
//	for idx, rule := range g {
//		c, i := count(rule)
//		if c == 1 {
//			answer[idx] = i
//		}
//
//	}
//}

//func count(m map[int]bool) (int, int) {
//	counter := 0
//	li := -1
//	for k, v := range m {
//		if v {
//			counter++
//			li = k
//		}
//	}
//	return counter, k
//}
func validNumber2(ticket []int64, v []rang, pm []int) bool {
	valid := true
	for idx, a := range ticket {
		if (a >= v[pm[idx]].min1 && a <= v[pm[idx]].max1) || (a >= v[pm[idx]].min2 && a <= v[pm[idx]].max2) {

		} else {
			valid = false
		}

	}
	return valid
}

func validNumber(a int64, m []rang) bool {
	valid := false
	for _, v := range m {
		if (a >= v.min1 && a <= v.max1) || (a >= v.min2 && a <= v.max2) {
			valid = true
		}
	}
	return valid
}

func splitRanges(s []string) rang {
	parts1 := strings.Split(s[1], "-")
	parts2 := strings.Split(s[3], "-")

	min1, _ := strconv.ParseInt(parts1[0], 10, 64)
	max1, _ := strconv.ParseInt(parts1[1], 10, 64)
	min2, _ := strconv.ParseInt(parts2[0], 10, 64)
	max2, _ := strconv.ParseInt(parts2[1], 10, 64)

	return rang{
		min1, max1, min2, max2,
	}
}

func main() {
	start := time.Now()
	solve2()
	fmt.Println(time.Now().Sub(start))

}
