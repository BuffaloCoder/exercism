package space

// Planet is type of string defining the name of a planet
type Planet string

const secondsInEarthOrbit = 31557600

var planetOrbitsComparedToEarth = map[Planet]float64{
	"Mercury": 0.2408467,
	"Venus":   0.61519726,
	"Earth":   1.0,
	"Mars":    1.8808158,
	"Jupiter": 11.862615,
	"Saturn":  29.447498,
	"Uranus":  84.016846,
	"Neptune": 164.79132,
}

// Age retuns the age in terms of orbital periods on a given
// planet with a specified age in seconds
func Age(seconds float64, planet Planet) float64 {
	var orbitComparedToEarth = planetOrbitsComparedToEarth[planet]
	var earthAge = seconds / secondsInEarthOrbit
	var givenPlanetAge = earthAge / orbitComparedToEarth

	return givenPlanetAge
}
