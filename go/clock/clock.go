package clock

import "fmt"

// Clock defines the underlying data needed to maintain a clock
type Clock struct {
	m int
}

// correctMinutes takes minutes and conforms it to the limits of the clock and
// wraps if it's needed
func correctMinutes(minutes int) int {
	minutes = minutes % 1440
	if minutes < 0 {
		minutes = 1440 + minutes
	}
	return minutes
}

// New returns a Clock set to the given hours and minutes. The time will wrap
// when appropriate
func New(hours, minutes int) Clock {
	c := Clock{correctMinutes(60*hours + minutes)}
	return c
}

// Add takes minutes and adds them to a Clock, returning the Clock in the
// process. The time will wrap when appropriate.
func (c Clock) Add(minutes int) Clock {
	c.m = correctMinutes(c.m + minutes)
	return c
}

// Subtract takes minutes and subtracts them from a Clock, returning the Clock in the
// process. The time will wrap when appropriate.
func (c Clock) Subtract(minutes int) Clock {
	c.m = correctMinutes(c.m - minutes)
	return c
}

// String tranforms a Clock struct into a string representation of time, of
// the format HH:MM.
func (c Clock) String() string {
	hours := c.m / 60
	minutes := c.m % 60

	return fmt.Sprintf("%02d:%02d", hours, minutes)
}
