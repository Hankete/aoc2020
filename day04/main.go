package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isPassportValid(passport map[string]string) bool {
	return validateByr(passport) && validateEyr(passport) && validateIyr(passport) &&
		validateHeight(passport) && validateHairColor(passport) && validateEyeColor(passport) &&
		validatePID(passport)
}

func validateByr(passport map[string]string) bool {
	if byrStr, ok := passport["byr"]; ok {
		byr, _ := strconv.ParseInt(byrStr, 10, 64)
		return byr >= 1920 && byr <= 2002
	}
	return false
}

func validateIyr(passport map[string]string) (answer bool) {
	if iyrStr, ok := passport["iyr"]; ok {
		iyr, _ := strconv.ParseInt(iyrStr, 10, 64)
		return iyr >= 2010 && iyr <= 2020
	}
	return false
}

func validateEyr(passport map[string]string) (answer bool) {
	if eyrStr, ok := passport["eyr"]; ok {
		eyr, _ := strconv.ParseInt(eyrStr, 10, 64)
		return eyr >= 2020 && eyr <= 2030
	}
	return false
}

func validateHeight(passport map[string]string) (answer bool) {
	hgtStr, ok := passport["hgt"]
	if !ok {
		return ok
	}
	if strings.HasSuffix(hgtStr, "cm") {
		hgtStr = strings.TrimSuffix(hgtStr, "cm")
		hgt, _ := strconv.ParseInt(hgtStr, 10, 64)
		return hgt >= 150 && hgt <= 193
	}
	if strings.HasSuffix(hgtStr, "in") {
		hgtStr = strings.TrimSuffix(hgtStr, "in")
		hgt, _ := strconv.ParseInt(hgtStr, 10, 64)
		return hgt >= 59 && hgt <= 76
	}
	return false
}

func validateHairColor(passport map[string]string) (answer bool) {
	hcl, ok := passport["hcl"]
	if !ok {
		return false
	}
	if hcl[0] != '#' || len(hcl) != 7 {
		return false
	}
	for i := 1; i < len(hcl); i++ {
		if (hcl[i] < '0' || hcl[i] > '9') && (hcl[i] < 'a' || hcl[i] > 'f') {
			return false
		}
	}
	return true
}

func validateEyeColor(passport map[string]string) (answer bool) {
	ecl, ok := passport["ecl"]
	if !ok {
		return false
	}
	if ecl != "amb" && ecl != "blu" && ecl != "brn" && ecl != "gry" && ecl != "grn" && ecl != "hzl" && ecl != "oth" {
		return false
	}
	return true
}

func validatePID(passport map[string]string) (answer bool) {
	pid, ok := passport["pid"]
	if !ok || len(pid) != 9 {
		return false
	}

	for _, val := range pid {
		if val < '0' || val > '9' {
			return false
		}
	}
	return true
}

func main() {
	f, err := os.Open("day04/day04-input.txt")
	if err != nil {
		panic("Can't open file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	lines := make([]string, 0)
	for scanner.Scan() {
		massStr := scanner.Text()
		lines = append(lines, massStr)
	}

	passports := make([]map[string]string, 0)
	var passport map[string]string
	i := 0
	passport = make(map[string]string)
	for _, line := range lines {
		if line == "" {
			i++
			passports = append(passports, passport)
			//fmt.Println(passports)
			passport = make(map[string]string)
		} else {
			pairs := strings.Split(line, " ")
			for _, pair := range pairs {
				split := strings.Split(pair, ":")
				if len(split) != 2 {
					fmt.Println(split)
				}
				passport[split[0]] = split[1]
				//fmt.Println(passport)
			}
		}
	}

	counter := 0
	for _, passport := range passports {
		if isPassportValid(passport) {
			counter++
		}
	}
	fmt.Println(counter, len(passports))
	fmt.Println(passports[len(passports)-1])
}
