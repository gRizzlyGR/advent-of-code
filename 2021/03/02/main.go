package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type rating []byte
type bitToRatings map[byte][]rating

type bitCriteriaFn func([]rating, []rating) []rating

type bitCriteria struct {
	pivot      byte
	pickRating bitCriteriaFn
}

var oxygenCriteria = bitCriteria{
	pivot:      '1',
	pickRating: mcv,
}

var co2Criteria = bitCriteria{
	pivot:      '0',
	pickRating: lcv,
}

func mcv(r1 []rating, r2 []rating) []rating {
	if len(r1) > len(r2) {
		return r1
	}

	return r2
}

func lcv(r1 []rating, r2 []rating) []rating {
	if len(r1) < len(r2) {
		return r1
	}

	return r2
}

func convertToDecimal(bit byte, i int, size int) int {
	return int(math.Pow(2, float64(size-i-1)) * float64(bit-'0'))
}

func pickRatings(b2r bitToRatings, criteria bitCriteria) []rating {
	ratingsByOne := b2r['1']
	ratingsByZero := b2r['0']

	// If quantity is the same, favor the reports that satisfy the bit criteria
	if len(ratingsByOne) == len(ratingsByZero) {
		return b2r[criteria.pivot]
	}

	return criteria.pickRating(ratingsByOne, ratingsByZero)
}

func findRating(ratings []rating, criteria bitCriteria) rating {
	for i := 0; len(ratings) != 1; i++ {

		b2r := make(bitToRatings)

		for _, r := range ratings {
			bit := r[i]
			b2r[bit] = append(b2r[bit], r)
		}

		ratings = pickRatings(b2r, criteria)
	}

	return ratings[0]
}

func computeLifeSupport(oxygen []byte, co2 []byte) int {
	var o int
	var c int
	size := len(oxygen)

	for i := 0; i < size; i++ {
		o += convertToDecimal(oxygen[i], i, size)
		c += convertToDecimal(co2[i], i, size)
	}

	return o * c
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	ratings := make([]rating, 0)

	for scanner.Scan() {
		rating := scanner.Text()
		ratings = append(ratings, []byte(rating))
	}

	oxygenRating := findRating(ratings, oxygenCriteria)
	co2Rating := findRating(ratings, co2Criteria)
	lifeSupport := computeLifeSupport(oxygenRating, co2Rating)

	fmt.Println(lifeSupport)
}
