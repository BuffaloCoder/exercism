// Package triangle maintains untilities for handling triangles
package triangle

import "math"

// Kind represents the different kinds of triangles as strings
type Kind string

const (
	// NaT represets not a triangle
	NaT = "Not a triangle"
	// Equ represents an equilateral triangle
	Equ = "Equialateral"
	// Iso represents an isosceles triangle
	Iso = "Isosceles"
	// Sca represents a scalene triangle
	Sca = "Scalene"
)

// isValidSide returns a bool indicating whether a given value could be
// a valid side of triangle
func isValidSide(side float64) bool {
	return side > 0 && !math.IsNaN(side) && !math.IsInf(side, 0)
}

// isValidTriangle determines if the given sides can form a valid triangle
func isValidTriangle(a, b, c float64) bool {
	return a+b >= c && b+c >= a && a+c >= b
}

// KindFromSides returns the type of triangle defined by it's side length, and
// if the side lengths indicate a triangle is not possible, returns a Kind
// that indicates as much
func KindFromSides(a, b, c float64) Kind {
	if !isValidSide(a) || !isValidSide(b) || !isValidSide(c) || !isValidTriangle(a, b, c) {
		// values given have at least one case where the sum of two sides is not
		// greater than the third; this is not a valid triangle
		return NaT
	}

	var equivalentSides = 0
	if a == b {
		equivalentSides++
	}
	if b == c {
		equivalentSides++
	}
	if a == c {
		equivalentSides++
	}

	switch equivalentSides {
	case 1:
		return Iso
	case 3:
		return Equ
	default:
		return Sca
	}
}
