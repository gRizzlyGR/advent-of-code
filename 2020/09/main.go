package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	preambleSize := 25

	nums := readNums()
	invalidNumber := findInvalidNumber(preambleSize, nums)
	goodNumbers := findNumbersThatSumToInvalidNumber(invalidNumber, nums)
	min, max := findLimits(goodNumbers)

	fmt.Println("Part 1:", invalidNumber)
	fmt.Println("Part 2:", min+max)
}

func readNums() []int {
	scanner := bufio.NewScanner(os.Stdin)

	nums := make([]int, 0)

	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, n)
	}

	return nums
}

func findInvalidNumber(preambleSize int, nums []int) int {
	for i := 0; i < len(nums)-preambleSize; i++ {
		preamble := nums[i : preambleSize+i]
		sums := buildSumsMap(preamble)
		n := nums[preambleSize+i]
		if _, ok := sums[n]; !ok {
			return n
		}
	}

	return -1
}

func buildSumsMap(preamble []int) map[int]struct{} {
	sums := make(map[int]struct{})
	size := len(preamble)
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			n1 := preamble[i]
			n2 := preamble[j]

			sums[n1+n2] = struct{}{}
		}
	}

	return sums
}

func findNumbersThatSumToInvalidNumber(invalidNumber int, nums []int) []int {
	goodNumbers := make([]int, 0)
	backup := invalidNumber

	j := 0
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		invalidNumber -= n
		goodNumbers = append(goodNumbers, n)

		// Restart
		if n == backup || invalidNumber < 0 {
			invalidNumber = backup
			goodNumbers = make([]int, 0)
			i = j
			j++
			continue
		} else if invalidNumber == 0 {
			break
		}
	}

	return goodNumbers
}

func findLimits(nums []int) (int, int) {
	sort.Ints(nums)

	return nums[0], nums[len(nums)-1]
}
