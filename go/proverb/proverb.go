// Package proverb contains utilities for generating proverds
package proverb

import "fmt"

// Proverb returns a proverd given a list of strings
func Proverb(rhyme []string) []string {
	var rhymeCount = len(rhyme)
	var proverb []string
	if rhymeCount > 0 {
		for i := 0; i < rhymeCount-1; i++ {
			proverb = append(proverb, fmt.Sprintf("For want of a %v the %v was lost.", rhyme[i], rhyme[i+1]))
		}
		proverb = append(proverb, fmt.Sprintf("And all for the want of a %v.", rhyme[0]))
	}

	return proverb
}
