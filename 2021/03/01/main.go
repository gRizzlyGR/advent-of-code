package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type counter map[byte]int

func initCounters(size int) []counter {
	counters := make([]counter, size)

	for i := 0; i < size; i++ {
		counters[i] = make(counter)
	}

	return counters
}

func buildRates(counters []counter) ([]byte, []byte) {
	size := len(counters)
	gamma := make([]byte, size)
	epsilon := make([]byte, size)

	for i, c := range counters {
		if c['0'] > c['1'] {
			gamma[i] = '0'
			epsilon[i] = '1'
		} else {
			gamma[i] = '1'
			epsilon[i] = '0'
		}
	}

	return gamma, epsilon
}

func convertToDecimal(bit byte, i int, size int) int {
	return int(math.Pow(2, float64(size-i-1)) * float64(bit-'0'))
}

func computeConsumption(gamma []byte, epsilon []byte) int {
	var g int
	var e int
	size := len(gamma)

	for i := 0; i < size; i++ {
		g += convertToDecimal(gamma[i], i, size)
		e += convertToDecimal(epsilon[i], i, size)
	}

	return g * e
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var counters []counter

	for scanner.Scan() {
		report := scanner.Bytes()

		if counters == nil {
			counters = initCounters(len(report))
		}

		for i, bit := range report {
			counters[i][bit]++
		}
	}

	gamma, epsilon := buildRates(counters)
	consumption := computeConsumption(gamma, epsilon)

	fmt.Println(consumption)
}
