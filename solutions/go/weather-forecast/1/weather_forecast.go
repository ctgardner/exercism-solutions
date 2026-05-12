// Package weather provides info about the weather in Goblinocus.
package weather

var (
	// CurrentCondition represents the current weather condition.
	CurrentCondition string
	// CurrentLocation represents the location of current weather conditions.
	CurrentLocation string
)

// Forecast sets the current weather condition in a city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
