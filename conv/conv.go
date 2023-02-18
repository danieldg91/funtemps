package conv

/*
  ALLE KONVERTERINGSFUNKSJONENE ER HER
*/

func FahrenheitToCelsius(value float64) float64 {
	celsius := (value - 32.0) * 5.0 / 9.0
	return celsius
}

func CelsiusToFahrenheit(value float64) float64 {
	fahrenheit := value*(9.0/5.0) + 32.0
	return fahrenheit
}

func FahrenheitToKelvin(value float64) float64 {
	kelvin := (value-32)*(5.0/9.0) + 273.15
	return kelvin
}

func KelvinToFahrenheit(value float64) float64 {
	fahrenheit := (value-273.15)*(9.0/5.0) + 32.0
	return fahrenheit
}

func CelsiusToKelvin(value float64) float64 {
	kelvin := value + 273.15
	return kelvin
}

func KelvinToCelsius(value float64) float64 {
	celsius := value - 273.15
	return celsius
}

// Du skal ikke formattere float64 i denne funksjonen
// Gjør formattering i main.go med fmt.Printf eller
// lignende
// return 56.67
// }
