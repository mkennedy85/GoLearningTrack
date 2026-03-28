// Package weather provides tools tp forecase weather.
package weather

var (
    // CurrentCondition stores the current weather.
	CurrentCondition string
    // CurrentLocation stores the current location.
	CurrentLocation  string
)

// Forecast returns the forecast for a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
