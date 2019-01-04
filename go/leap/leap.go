// Package leap contains tulities surrounding leap years
package leap

// IsLeapYear returns a bool indicating whether or not a given year is
// considered a leap year.
func IsLeapYear(year int) bool {
	var yearDivisibleBy4 = year%4 == 0
	var yearDivisibleBy100 = year%100 == 0
	var yearDivisibleBy400 = year%400 == 0
	// To be a leap year, the given year must be:
	//  divisble by four AND
	//   either not be divisble by 100 or be divisble by 400
	if yearDivisibleBy4 &&
		(!yearDivisibleBy100 || yearDivisibleBy400) {
		return true
	}
	return false
}
