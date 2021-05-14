package utils

// IsDigitPresent IsDigitPresent
func IsDigitPresent(number, d int) bool {
	// Breal loop if d is present as digit
	for number > 0 {
		if number%10 == d {
			break
		}

		number = number / 10
	}

	// If loop broke
	return (number > 0)
}

// IsDigitPresentUint64 IsDigitPresentUint64
func IsDigitPresentUint64(number uint64, d uint64) bool {
	// Breal loop if d is present as digit
	for number > 0 {
		if number%10 == d {
			break
		}

		number = number / 10
	}

	// If loop broke
	return (number > 0)
}
