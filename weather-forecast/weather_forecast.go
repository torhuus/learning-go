// Package weather defines the package we are using.
package weather

// CurrentCondition is a string variable.
var CurrentCondition string

// CurrentLocation is a string variable.
var CurrentLocation string

// Forecast function takes in a city and a condition and returns an appropriate string.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
