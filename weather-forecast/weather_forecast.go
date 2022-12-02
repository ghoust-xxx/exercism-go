// Package weather forecast the current weather condition of various cities
// in Goblinocus.
package weather

// CurrentCondition have current conditions.
var CurrentCondition string
// CurrentLocation have current location.
var CurrentLocation string

// Forecast returns forecast for city with condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
