package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	req := map[string]struct{}{
		"byr": {},
		"iyr": {},
		"eyr": {},
		"hgt": {},
		"hcl": {},
		"ecl": {},
		"pid": {},
	}

	nValidFields := 0
	nValidPass := 0
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())

		if len(fields) == 0 {
			if len(req) == nValidFields {
				nValidPass++
			}
			nValidFields = 0
		}

		for _, field := range fields {
			parts := strings.Split(field, ":")
			key := parts[0]
			val := parts[1]
			if _, ok := req[key]; ok {
				if validate(key, val) {
					nValidFields++
				}
			}
		}
	}

	fmt.Println(nValidPass)
}

func validate(key, val string) bool {
	switch key {
	case "byr":
		return checkByr(val)
	case "iyr":
		return checkIyr(val)
	case "eyr":
		return checkEyr(val)
	case "hgt":
		return checkHgt(val)
	case "hcl":
		return checkHcl(val)
	case "ecl":
		return checkEcl(val)
	case "pid":
		return checkPid(val)
	default:
		return false
	}

	return true
}

func checkByr(val string) bool {
	n, err := strconv.Atoi(val)
	if err != nil || n < 1920 || n > 2002 {
		return false
	}

	return true
}

func checkIyr(val string) bool {
	n, err := strconv.Atoi(val)
	if err != nil || n < 2010 || n > 2020 {
		return false
	}

	return true
}

func checkEyr(val string) bool {
	n, err := strconv.Atoi(val)
	if err != nil || n < 2020 || n > 2030 {
		return false
	}

	return true
}

func checkHcl(val string) bool {
	_, err := strconv.ParseInt(val[1:], 16, 64)
	if len(val) != 7 || val[0] != '#' || err != nil {
		return false
	}

	return true
}

func checkHgt(val string) bool {
	n, err := strconv.Atoi(val[:len(val)-2])
	measure := val[len(val)-2:]
	if err != nil {
		return false
	}

	switch measure {
	case "cm":
		if n < 150 || n > 193 {
			return false
		}
	case "in":
		if n < 59 || n > 76 {
			return false
		}
	default:
		return false
	}

	return true
}

func checkEcl(val string) bool {
	colors := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

	for _, color := range colors {
		if val == color {
			return true
		}
	}

	return false
}

func checkPid(val string) bool {
	match, _ := regexp.MatchString("^\\d{9}$", val)
	return match
}
