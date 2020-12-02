package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	valid := 0

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		limits := strings.Split(line[0], "-")

		min, _ := strconv.Atoi(limits[0])
		max, _ := strconv.Atoi(limits[1])

		letter := string(line[1][0])
		password := line[2]

		count := strings.Count(password, letter)

		if count >= min && count <= max {
			valid++
		} 
	}

	fmt.Println(valid)
}
