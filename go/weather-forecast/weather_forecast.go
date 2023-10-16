// Package weather provides functionality for weather forecasting.
package weather

// CurrentCondition represents the current weather condition.
var CurrentCondition string
// CurrentLocation represents the current location for which the weather is being forecasted.
var CurrentLocation string

// Forecast returns a string representing the forecast for the given city and condition.
// It updates the CurrentLocation and CurrentCondition variables and combines them with.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
