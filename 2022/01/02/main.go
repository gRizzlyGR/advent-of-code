package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	allCalories := make([]int, 0)

	calories := 0

	for scanner.Scan() {
		line := scanner.Text()

		currentCalories, _ := strconv.Atoi(line)
		calories += currentCalories

		if line == "" {
			allCalories = append(allCalories, calories)
			calories = 0
			continue
		}
	}

	allCalories = append(allCalories, calories)

	sort.Ints(allCalories)

	topCalories := 0

	for _, top := range allCalories[len(allCalories)-3:] {
		topCalories += top
	}

	fmt.Println(topCalories)
}
