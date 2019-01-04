// Package scrabble contains utilities for analyzing scrabble games
package scrabble

import "strings"

var letterScores = map[rune]int{}

// Score takes a string representing a scrabble word and returns the score
// for the word (does not check for validity of the word itself)
func Score(word string) int {
	word = strings.ToLower(word)
	var score int
	for _, char := range word {
		switch char {
		case 'a', 'e', 'i', 'l', 'n', 'o', 'r', 's', 't', 'u':
			score++
		case 'd', 'g':
			score += 2
		case 'b', 'c', 'm', 'p':
			score += 3
		case 'f', 'h', 'v', 'w', 'y':
			score += 4
		case 'k':
			score += 5
		case 'j', 'x':
			score += 8
		case 'q', 'z':
			score += 10
		}
	}

	return score
}
