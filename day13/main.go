package main

import (
	"AOC2020/ioutils"
	"fmt"
	"github.com/deanveloper/modmath"
	"math"
	"strconv"
	"strings"
)

func solve1() {
	path := "day13/day13-test.txt"
	lines := ioutils.ReadLines(path)

	ids := strings.Split(lines[1], ",")
	arrTime, _ := strconv.ParseInt(lines[0], 10, 64)
	times := make([]int64, 0)

	currWait := int64(math.MaxInt32)
	currId := int64(math.MaxInt32)

	for _, id := range ids {
		if id == "x" {
			continue
		}
		depTime, _ := strconv.ParseInt(id, 10, 64)
		times = append(times, depTime)

		rest := arrTime % depTime
		//x := arrTime / depTime
		//y := depTime * x
		wait := depTime - rest

		if wait < currWait {
			currWait = wait
			currId = depTime
		}
	}

	fmt.Println(currWait, currId, currWait*currId)
}

//19

func solve3() {
	path := "day13/day13-input.txt"
	lines := ioutils.ReadLines(path)

	ids := strings.Split(lines[1], ",")

	varCounter := 0
	for _, id := range ids {
		if id != "x" {
			varCounter += 1
		}
	}

	A := make([]uint64, varCounter)
	R := make([]uint64, varCounter)
	CRT := make([]modmath.CrtEntry, 0)
	i := -1
	for idx, id := range ids {
		if id == "x" {
			continue
		}
		i++
		val, _ := strconv.ParseUint(id, 10, 64)
		A[i] = val
		R[i] = (val - uint64(idx)%val) % val
		fmt.Println(A[i], R[i], idx)
		cr := modmath.CrtEntry{
			A: int(A[i]),
			M: int(R[i]),
		}
		CRT = append(CRT, cr)

	}

	//n := modmath.SolveCrtMany(CRT)

	fmt.Println(R)
	fmt.Println(A)

}
func main() {
	solve3()
}
