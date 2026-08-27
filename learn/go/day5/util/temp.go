package util

// ToFahrenheit converts Celsius to Fahrenheit.
func ToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

// ToCelsius converts Fahrenheit to Celsius.
func ToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}
