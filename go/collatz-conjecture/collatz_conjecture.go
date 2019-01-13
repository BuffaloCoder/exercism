package collatzconjecture

import "errors"

// CollatzConjecture takes a given number, performs the Collatz Conjecture
// algorithm on it, and returns the number of steps it takes to get to 1
func CollatzConjecture(num int) (int, error) {
	if num < 1 {
		return 0, errors.New("Input must be greater than 0")
	}

	var steps int
	// While not proven, it's currently believed that any value will eventually
	// reach 1
	for num != 1 {
		steps++
		// We could use % and / to determine if it's divisible and then divide by
		// 2, but the bitwise operations are generally faster (the collatz conjecture
		// can lead to large step counts, so might as well shave time when possible)
		if (num & 1) == 0 {
			num = num >> 1
		} else {
			num = 3*num + 1
		}
	}

	return steps, nil
}
