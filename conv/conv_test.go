package conv

import (
	"math"
	"testing"
)

// JANIS' FUNCTION "withinTolerance"

func withinTolerance(a, b, error float64) bool {
	// Først sjekk om tallene er nøyaktig like
	if a == b {
		return true
	}

	difference := math.Abs(a - b)

	// Siden vi skal dele med b, må vi sjekke om den er 0
	// Hvis b er 0, returner avgjørelsen om d er mindre enn feilmarginen
	// som vi aksepterer
	if b == 0 {
		return difference < error
	}

	// Tilslutt sjekk den relative differanse mot feilmargin
	return (difference / math.Abs(b)) < error
}

/*  FIRST TEST = FAHRENHEIT TO CELCIUS   */
func TestFahrenheitToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 134, want: 56.67},
	}

	for _, tc := range tests {
		got := FahrenheitToCelsius(tc.input)
		if !withinTolerance(tc.want, got, 1e-3) {
			t.Errorf("want: %12.2f, got: %12.2f", tc.want, got)
		}
	}
}

/*  SECOND TEST = CELCIUS TO FAHRENHEIT */
func TestCelsiusToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 0.0, want: 32.0},
	}
	for _, tc := range tests {
		got := CelsiusToFahrenheit(tc.input)
		if !withinTolerance(tc.want, got, 1e-3) {
			t.Errorf("want: %12.2f; got %12.2f", tc.want, got)
		}
	}
}

/*  THIRD TEST = FAHRENHEIT TO KELVIN */
func TestFahrenheitToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 32, want: 273.15},
	}

	for _, tc := range tests {
		got := FahrenheitToKelvin(tc.input)
		if !withinTolerance(tc.want, got, 1e-3) {
			t.Errorf("want: %12.2f, got: %12.2f", tc.want, got)
		}
	}
}

/*  FOURTH TEST = KELVIN TO FAHRENHEIT */
func TestKelvinToFahrenheit(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 273.15, want: 32},
	}

	for _, tc := range tests {
		got := KelvinToFahrenheit(tc.input)
		if !withinTolerance(tc.want, got, 1e-3) {
			t.Errorf("want: %12.2f, got: %12.2f", tc.want, got)
		}
	}
}

/*  FIFTH TEST FUNCTION = CELSIUS TO KELVIN */
func TestCelsiusToKelvin(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 0, want: 273.15},
	}

	for _, tc := range tests {
		got := CelsiusToKelvin(tc.input)
		if !withinTolerance(tc.want, got, 1e-3) {
			t.Errorf("want: %12.2f, got: %12.2f", tc.want, got)
		}
	}
}

/*  SIXTH TEST FUNCTION = KELVIN TO CELSIUS */
func TestKelvinToCelsius(t *testing.T) {
	type test struct {
		input float64
		want  float64
	}

	tests := []test{
		{input: 273.15, want: 0},
	}

	for _, tc := range tests {
		got := KelvinToCelsius(tc.input)
		if !withinTolerance(tc.want, got, 1e-3) {
			t.Errorf("want: %12.2f, got: %12.2f", tc.want, got)
		}
	}
}
