package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	target := 2020

	product := -1
	differences := make(map[int]int)

	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		differences[n] = target - n
	}

	for _, diff := range differences {
		if n, ok := differences[diff]; ok {
			fmt.Println(n, diff)
			product = n * diff
			break
		}
	}

	fmt.Println(product)
}
