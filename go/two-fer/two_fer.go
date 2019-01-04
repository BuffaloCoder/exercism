// Package twofer implements a simple library for outputing two for one exachanges.
package twofer

// ShareWith returns a string indicating who gets how many of something. If no
// name is given, the string will use "you"
func ShareWith(name string) string {
	if len(name) == 0 {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
