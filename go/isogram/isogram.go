// Package isogram contains utilities for handling isograms
package isogram

import "unicode"

// IsIsogram takes a word and will return a boolean indicating whether or not
// it is an isogram
func IsIsogram(word string) bool {
	var letterMap = map[rune]bool{}

	for _, letter := range word {
		letter = unicode.ToLower(letter)
		if letter == ' ' || letter == '-' {
			continue
		} else if letterMap[letter] {
			return false
		} else {
			letterMap[letter] = true
		}
	}

	return true
}
