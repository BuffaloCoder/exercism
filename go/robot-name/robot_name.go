package robotname

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const maxCount int = 26 * 26 * 10 * 10 * 10
const letterOffset rune = 'A'

const firstLetterBase int = 26000
const secondLetterBase int = 1000

// maintains our list of available names
var availNameIDs []int
var availIDCount = maxCount

// Robot represents a robot with a given name
type Robot struct {
	name string
}

func init() {
	availNameIDs = make([]int, maxCount)
	for i := 0; i < maxCount; i++ {
		availNameIDs[i] = i
	}
}

func getNameFromID(id int) string {
	// Generate the letters, which are determined by the offset from their
	// starting (base) positions. Ex. if the id is 26000, that means theres a
	// "1" in the firstLetter place, so you add a one to A and get B (but the
	// rest will be the starting "0" positions)
	// Note: the idea for this way of getting a name came from the example code
	// on the github page (though re-implemented)
	firstLetter := letterOffset + rune(id/firstLetterBase)
	remainder := id % firstLetterBase
	secondLetter := letterOffset + rune(remainder/secondLetterBase)
	number := remainder % secondLetterBase

	return fmt.Sprintf("%s%s%03d", string(firstLetter), string(secondLetter), number)
}

// Name will either return the robot's name if it exists, or will generate a
// new, random name of type AA### that does not match an existing robot name.
func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	if availIDCount == 0 {
		return "", errors.New("All names haves been requested")
	}

	// Select a random id from the remaining ids
	rand.Seed(time.Now().UnixNano())
	idIndex := rand.Intn(availIDCount)
	r.name = getNameFromID(availNameIDs[idIndex])

	// "Remove" the last taken id by swapping it with the current last position and
	// decrementing the counter. This way we don't have to worry about reallocation.
	availNameIDs[idIndex], availNameIDs[availIDCount-1] = availNameIDs[availIDCount-1], availNameIDs[idIndex]
	availIDCount--
	return r.name, nil
}

// Reset will revert the Robot's name to the empty state
func (r *Robot) Reset() {
	r.name = ""
}
