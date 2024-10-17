import "math"

func maximumSwap(num int) int {
	if num < 12 {
		return num
	}

	n := num

	numLen := int(math.Log10(float64(num))) + 1

	convertedToSlice := make([]int, numLen)

	digits := make([]int, 10)
	lastDigit := 0

	for i := 0; i < numLen; i++ {
		lastDigit = n % 10
		n /= 10

		convertedToSlice[numLen-i-1] = lastDigit

		if digits[lastDigit] == 0 {
			digits[lastDigit] = numLen - i - 1
		}
	}

	for i := 0; i < numLen; i++ {
		for j := 9; j > convertedToSlice[i]; j-- {
			if digits[j] > i {
				convertedToSlice[i], convertedToSlice[digits[j]] = convertedToSlice[digits[j]], convertedToSlice[i]

				result := 0

				for k := 0; k < numLen; k++ {
					result = result*10 + convertedToSlice[k]
				}

				return result
			}
		}
	}

	return num
}
