// Package luhn contains utilities for validating numbers via Luhn's algorithm
package luhn

import (
	"errors"
	"unicode"
)

func validateAndParseString(toParse string) ([]int, error) {
	var numList []int

	// Strip out spaces. If any non-number characters other than spaces are
	// found, then this is not a valid number.
	for _, char := range toParse {
		if unicode.IsNumber(char) {
			numList = append(numList, int(char-'0'))
		} else if char != ' ' {
			return numList, errors.New("Input contains non-numeric characters other than strings")
		}
	}

	if len(numList) < 2 {
		return numList, errors.New("Input contains only one digit")
	}

	return numList, nil
}

// Valid takes a given string and determines if it's valid according to
// Luhn's algorithm
func Valid(toValidate string) bool {
	var luhnsNumber int

	numList, err := validateAndParseString(toValidate)
	if err != nil {
		return false
	}

	var isSecondDigit bool
	for i := len(numList) - 1; i >= 0; i-- {
		num := numList[i]

		if isSecondDigit {
			if num *= 2; num > 9 {
				num -= 9
			}
		}
		luhnsNumber += num

		// setup isSecondDigit for the next number
		isSecondDigit = !isSecondDigit
	}

	return luhnsNumber%10 == 0
}
