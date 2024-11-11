// Package weather provides information about the weather.
package weather

// CurrentCondition refers to the status of the weather.
var CurrentCondition string

// CurrentLocation refers to the location of the given weather.
var CurrentLocation string

// Forecast returns a string with the condition in a given location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
