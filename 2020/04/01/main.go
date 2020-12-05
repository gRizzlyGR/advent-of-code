package main

import (
	"bufio"
	"fmt"
	"os"
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
			key := strings.Split(field, ":")[0]
			if _, ok := req[key]; ok {
				nValidFields++
			}
		}
	}

	fmt.Println(nValidPass)
}
