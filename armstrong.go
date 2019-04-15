package armstrong

import (
	"math/big"
)

func IsArmstrong(number *big.Int) bool {

	digits := getDigitsFromNumber(number)
	sumOfPowers := sumOfPowers(digits)
	return equals(number, sumOfPowers)
}

func getDigitsFromNumber(number *big.Int) []int64 {

	digits := make([]int64, 0)
	ten := big.NewInt(10)
	auxNumber := new(big.Int).Set(number)

	for greaterThanZero(auxNumber) {

		remainder := new(big.Int).Mod(auxNumber, ten).Int64()
		digits = append(digits, remainder)

		auxNumber.Div(auxNumber, ten)
	}

	return digits // in reverse order, but it does not affect the result
}

func sumOfPowers(digits []int64) *big.Int {

	numberOfDigits := big.NewInt(int64(len(digits)))
	sum := big.NewInt(0)

	for _, digit := range digits {
		power := new(big.Int).Exp(big.NewInt(digit), numberOfDigits, nil)
		sum.Add(sum, power)
	}

	return sum
}

func greaterThanZero(number *big.Int) bool {
	return number.Cmp(big.NewInt(0)) > 0
}

func equals(number *big.Int, secondNumber *big.Int) bool {
	return number.Cmp(secondNumber) == 0
}
