package spaceage

type Planet string

var year = 31557600.0
var orbitalPeriod = map[Planet]float64{
	"Mercury": year * 0.2408467,
	"Venus":   year * 0.61519726,
	"Earth":   year,
	"Mars":    year * 1.8808158,
	"Jupiter": year * 11.862615,
	"Saturn":  year * 29.447498,
	"Uranus":  year * 84.016846,
	"Neptune": year * 164.79132,
}

func Age(seconds float64, planet Planet) float64 {
	if period, ok := orbitalPeriod[planet]; ok {
		return seconds / period
	}
	return -1
}
