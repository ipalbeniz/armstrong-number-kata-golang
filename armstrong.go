package armstrong

import (
	"github.com/thoas/go-funk"
	"math"
)

func IsArmstrong(number float64) bool {
	digits := getDigitsFromNumber(number)
	return number == sumOfPowers(digits)
}

func getDigitsFromNumber(number float64) []int {
	digits := make([]int, 0)

	for exponentOfTen := 0;; exponentOfTen ++ {
		powerOfTen := math.Pow(10, float64(exponentOfTen))
		if powerOfTen > number {
			break
		}
		digit := int(math.Mod(number / powerOfTen, 10))
		digits = append(digits, digit)
	}

	return funk.ReverseInt(digits)
}

func sumOfPowers(digits []int) float64 {
	powers := funk.Map(digits, func(digit int) float64 {
		return math.Pow(float64(digit), float64(len(digits)))
	})

	return funk.Reduce(powers, func(sum float64, power float64) float64 {
		return sum + power
	}, 0)
}
