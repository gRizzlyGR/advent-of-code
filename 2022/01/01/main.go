package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	maxCalories := -1
	calories := 0

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if calories > maxCalories {
				maxCalories = calories
			}

			calories = 0
			continue
		}

		currentCalories, _ := strconv.Atoi(line)
		calories += currentCalories
	}

	fmt.Println(maxCalories)
}
