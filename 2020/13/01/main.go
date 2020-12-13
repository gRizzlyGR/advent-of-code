package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	myDeparture, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	buses := readBuses(scanner.Text())

	bus, timeGap := findBest(myDeparture, buses)
	fmt.Println(bus * timeGap)
}

func readBuses(s string) []int {
	buses := make([]int, 0)

	split := strings.Split(s, ",")

	for _, c := range split {
		if c != "x" {
			n, _ := strconv.Atoi(c)
			buses = append(buses, n)
		}
	}

	return buses
}

func findBest(myDeparture int, buses []int) (int, int) {
	earliest := math.MaxInt64
	firstAvailable := 0

	for _, bus := range buses {
		tmp := float64(myDeparture) / float64(bus)
		busDeparture := bus * int(math.Ceil(tmp))

		if busDeparture < earliest {
			earliest = busDeparture
			firstAvailable = bus
		}
	}

	timeGap := earliest - myDeparture

	return firstAvailable, timeGap
}
