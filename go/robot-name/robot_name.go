package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// maintains our list of consumed names
var robotNames = map[string]bool{
	"": true,
}
var usedNameCount int

const letterOffset rune = 'A'

// Robot represents a robot with a given name
type Robot struct {
	name string
}

// Name will either return the robot's name if it exists, or will generate a
// new, random name of type AA### that does not match an existing robot name.
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	if usedNameCount == 26*26*10*10*10 {
		return "", errors.New("All names haves been requested")
	}
	// If it has, then we should increment the name to the next available name.
	// This is done to avoid a possible near-infinite loop with generating continously
	// random names, and is relatively random since collisions should not be common and
	// the base string is random
	for robotNames[r.name] {
		rand.Seed(time.Now().UnixNano())
		firstLetter := letterOffset + rune(rand.Intn(26))
		secondLetter := letterOffset + rune(rand.Intn(26))
		number := rand.Intn(1000)

		r.name = fmt.Sprintf("%s%s%03d", string(firstLetter), string(secondLetter), number)
	}

	robotNames[r.name] = true
	usedNameCount++
	if usedNameCount%1000 == 0 {
		fmt.Println(usedNameCount)
	}
	return r.name, nil
}

// Reset will revert the Robot's name to the empty state
func (r *Robot) Reset() {
	r.name = ""
}
