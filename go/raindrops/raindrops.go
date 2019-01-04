// Package raindrops is a set of utilities for managing drop-speak
package raindrops

import "strconv"

// Convert takes an integer and converts it to drop-speak
func Convert(drop int) string {
	var dropSpeak string

	// Check the drop-speak applicable factors
	if drop%3 == 0 {
		dropSpeak += "Pling"
	}
	if drop%5 == 0 {
		dropSpeak += "Plang"
	}
	if drop%7 == 0 {
		dropSpeak += "Plong"
	}

	// If conversion is not applicable, set it to the number
	if dropSpeak == "" {
		dropSpeak = strconv.Itoa(drop)
	}

	return dropSpeak
}
