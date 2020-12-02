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

		to, _ := strconv.Atoi(limits[0])
		from, _ := strconv.Atoi(limits[1])

		to--
		from--

		letter := line[1][0]
		password := line[2]

		if (letter == password[to] && letter != password[from]) || (letter != password[to] && letter == password[from]) {
			valid++
		}
	}

	fmt.Println(valid)
}
