package grains

import "errors"

// Square takes a number representing the a square on teh chessboard from 1-64
// and returns the number of grains on it.
func Square(square int) (uint64, error) {
	if square < 1 {
		return 0, errors.New("The squares begin at 1; any number below that is invalid")
	}
	if square > 64 {
		return 0, errors.New("The squares end at 64; any number above that is invalid")
	}

	// In this case, we're only dealing with powers of two, which can be handled by shifting
	// the bit to the left by one for each square.
	return 1 << uint(square-1), nil
}

// Total returns the total count of all squares on the board added together
func Total() uint64 {
	// The total is just going to be 2^64-1
	return 1<<64 - 1
}
