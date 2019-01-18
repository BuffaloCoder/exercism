package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// maintains our list of consumed names
var robotNames = map[string]bool{}
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
	rand.Seed(time.Now().UnixNano())
	firstLetter := letterOffset + rune(rand.Intn(26))
	secondLetter := letterOffset + rune(rand.Intn(26))
	number := rand.Intn(1000)

	initialName := fmt.Sprintf("%s%s%03d", string(firstLetter), string(secondLetter), number)
	name := initialName
	// If it has, then we should increment the name to the next available name.
	// This is done to avoid a possible near-infinite loop with generating continously
	// random names, and is relatively random since collisions should not be common and
	// the base string is random
	for robotNames[name] {
		if number != 999 {
			number++
		} else if firstLetter != 'Z' {
			number = 0
			firstLetter++
		} else if secondLetter != 'Z' {
			number = 0
			firstLetter = 'A'
			secondLetter++
		} else {
			number = 0
			firstLetter = 'A'
			secondLetter = 'A'
		}

		name = fmt.Sprintf("%s%s%03d", string(firstLetter), string(secondLetter), number)
		if name == initialName {
			return "", errors.New("All names haves been requested")
		}
	}

	robotNames[name] = true
	usedNameCount++
	if usedNameCount % 1000 {
		fmt.Println(usedNameCount)
	}
	r.name = name
	return r.name, nil
}

// Reset will revert the Robot's name to the empty state
func (r *Robot) Reset() {
	r.name = ""
}
