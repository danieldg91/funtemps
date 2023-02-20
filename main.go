package main

import (
	"danieldg91/funtemps/conv"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

var input, output string
var cels, kelv, fahr float64

func init() {

	// definition of flags -F, -K, -C and -out
	flag.Float64Var(&fahr, "F", 0.0, "temperatur i grader fahrenheit")
	flag.Float64Var(&cels, "C", 0.0, "temperatur i grader celsius")
	flag.Float64Var(&kelv, "K", 0.0, "temperatur i grader kelvin")
	flag.StringVar(&output, "out", "", "Calculates output-temperatures in C - Celsius, K - Kelvin, F - Fahrenheit)")
	flag.StringVar(&input, "in", "", "Enter input '-C' for Celsius, '-K' for Kelvin or '-F' for Fahrenheit)")
}

func main() {

	flag.Parse()

	// Conversion of temperatures
	var result float64
	var input float64
	var out string

	// -F flag
	if fahr != 0 {
		input = fahr
		if output == "F" {
			fmt.Println("Error: '-out F'. No conversion between identical temperature scales")
		} else if output == "C" {
			result = conv.FahrenheitToCelsius(fahr)
			out = " °C"
		} else if output == "K" {
			result = conv.FahrenheitToKelvin(fahr)
			out = " °K"
		} else {
			fmt.Println("Invalid output flag. Choose either -C or -K")
		}
		// -C flag
	} else if cels != 0 {
		input = cels
		if output == "C" {
			fmt.Println("Error: '-out C'. No conversion between identical temperature scales")
		} else if output == "F" {
			result = conv.CelsiusToFahrenheit(cels)
			out = " °F"
		} else if output == "K" {
			result = conv.CelsiusToKelvin(cels)
			out = " °K"
		} else {
			fmt.Println("Invalid output flag. Choose either -F or -K")
		}

	} else if kelv != 0 {
		input = kelv
		if output == "K" {
			fmt.Println("Error: no conversion between identical temperature scales")
		} else if output == "F" {
			result = conv.KelvinToFahrenheit(kelv)
			out = " °F"
		} else if output == "C" {
			result = conv.KelvinToCelsius(kelv)
			out = " °C"
		} else {
			fmt.Println("Invalid output flag. Choose either -C -F")
		}
	} else {
		fmt.Println("You must enter a valid input temperature scale (-C, -F or -K).")
	}

	// Formatterer input etter Janis' spesifikasjoner.
	formattedInput := strconv.FormatFloat(input, 'f', -1, 64)
	if strings.Contains(formattedInput, ".") {
		formattedInput = strings.TrimRight(formattedInput, "0")
		if formattedInput[len(formattedInput)-1] == '.' {
			formattedInput = formattedInput[:len(formattedInput)-1]
		}
	}
	// Formatterer output med funksjon lengre nede
	formattedResult := formatNumber(result)

	fmt.Printf("%14s%s = %14s%s\n", formattedInput, getTemperatureScale(cels, fahr, kelv), formattedResult, out)
	// må være %s for at format-funksjonene skal fungere.

}

// Når denne blir kalt i main, så tar den imot verdiene fra cels, fahr og kelv og returnerer rett bokstav på slutten av float64 tallet.
func getTemperatureScale(c, f, k float64) string {
	if c != 0 {
		return " °C"
	} else if f != 0 {
		return " °F"
	} else if k != 0 {
		return " °K"
	}
	return ""
}

func formatNumber(num float64) string {
	// sjekker om det er desimaler...
	if num == float64(int(num)) {
		// hvis ikke, returnerer integer delen uten desimaler
		return strconv.Itoa(int(num))
	}

	// Float til string med 2 desimaler
	str := strconv.FormatFloat(num, 'f', 2, 64)

	parts := strings.Split(str, ".")
	intPart := parts[0]
	fracPart := ""
	if len(parts) == 2 {
		fracPart = "." + parts[1]
	}

	// Setter inn mellomrom for hver tredje siffer.
	n := len(intPart)
	if n > 3 {
		for i := n - 3; i > 0; i -= 3 {
			intPart = intPart[:i] + "'" + intPart[i:]
		}
	}

	return intPart + fracPart
}
