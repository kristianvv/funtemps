package conv

// Konverterer Farenheit ti Celsius
func FahrenheitToCelsius(value float64) float64 {
	return (value - 32.0) * (5.0 / 9.0)
}

// Konverterer Celsius til Farenheit
func CelsiusToFahrenheit(value float64) float64 {
	return (value * 9 / 5) + 32
}

// Konverterer Farenheit ti Celsius
func KelvinToFahrenheit(value float64) float64 {
	return (value-273.15)*9/5 + 32
}

// Konverterer Farenheit ti Celsius
func FahrenheitToKelvin(value float64) float64 {
	return (value-32)*5/9 + 273.15
}

// Konverterer Farenheit ti Celsius
func CelsiusToKelvin(value float64) float64 {
	return (value + 273.15)
}

// Konverterer Farenheit ti Celsius
func KelvinToCelsius(value float64) float64 {
	return (value - 273.15)
}
