// Package space calculates someone's age according to planet age.
package space

// Planet is a string.
type Planet string

const earth = 31557600

var planetAge = map[Planet]float64{
	"Earth":   earth,
	"Mercury": earth * 0.2408467,
	"Venus":   earth * 0.61519726,
	"Mars":    earth * 1.8808158,
	"Jupiter": earth * 11.862615,
	"Saturn":  earth * 29.447498,
	"Uranus":  earth * 84.016846,
	"Neptune": earth * 164.79132,
}

// Age calculates someone's age.
func Age(seconds float64, planet Planet) float64 {
	return seconds / planetAge[planet]
}
