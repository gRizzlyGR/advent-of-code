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
	nums := make([]int, 0)
	N1ToN2PlusN3 := make(map[int]int)
	N2PlusN3ToN1 := make(map[int]int)

	for scanner.Scan() {
		n1, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, n1)
		n2PlusN3 := target - n1
		N1ToN2PlusN3[n1] = n2PlusN3
		N2PlusN3ToN1[n2PlusN3] = n1
	}

	for n2PlusN3, n1 := range N2PlusN3ToN1 {
		for _, n := range nums {
			n3 := n2PlusN3 - n
			n1PlusN2 := N1ToN2PlusN3[n3]

			n2 := n1PlusN2 - n1

			_, ok1 := N1ToN2PlusN3[n1]
			_, ok2 := N1ToN2PlusN3[n2]
			_, ok3 := N1ToN2PlusN3[n3]

			if ok1 && ok2 && ok3 {
				fmt.Println(n1, n2, n3)
				product = n1 * n2 * n3
				fmt.Println(product)
				return
			}
		}
	}
}