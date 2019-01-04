// Package hamming contains a set of utilities for determining the hamming
// distance between two DNA strands
package hamming

import "errors"

// Distance calculates the hamming distance between two strings
// representing DNA strands
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("DNA strands are not equal in length and cannot be compared")
	}
	var distance = 0
	// only need to iterate over the one since they are confirmed to be of the same length
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}
