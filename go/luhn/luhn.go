// Package luhn contains utilities for validating inputs via Luhn's algorithm
package luhn

import "unicode"

// Valid takes a string input and evaluates to see if it passes
// Luhn's verification algorithm
func Valid(toValidate string) bool {
	var luhnsNumber int
	var numberCount int

	for i := len(toValidate) - 1; i >= 0; i-- {
		char := rune(toValidate[i])
		// If the char is anything but a number or a space, fail the assertion.
		// Otherwise, if it's a space, then skip to the next char.
		if unicode.IsSpace(char) {
			continue
		}
		if !unicode.IsNumber(char) {
			return false
		}

		num := int(char - '0')
		numberCount++

		if numberCount%2 == 0 {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}

		luhnsNumber += num
	}

	if numberCount < 2 {
		return false
	}

	return luhnsNumber%10 == 0
}
