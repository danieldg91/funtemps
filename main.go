package main

import (
	"flag"
	"fmt"

	"danieldg91/funtemps/conv"
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

	// -F flag
	if fahr != 0 {
		input = fahr
		if output == "F" {
			fmt.Println("Error: '-out F'. No conversion between identical temperature scales")
		} else if output == "C" {
			result = conv.FahrenheitToCelsius(fahr)
		} else if output == "K" {
			result = conv.FahrenheitToKelvin(fahr)
		} else {
			fmt.Println("Invalid output flag or value '0'. Choose either -C, -F or -K")
		}
		// -C flag
	} else if cels != 0 {
		input = cels
		if output == "C" {
			fmt.Println("Error: '-out C'. No conversion between identical temperature scales")
		} else if output == "F" {
			result = conv.CelsiusToFahrenheit(cels)
		} else if output == "K" {
			result = conv.CelsiusToKelvin(cels)
		} else {
			fmt.Println("Invalid output flag or value '0'. Choose either -C, -F or -K")
		}

	} else if kelv != 0 {
		input = kelv
		if output == "K" {
			fmt.Println("Error: no conversion between identical temperature scales")
		} else if output == "F" {
			result = conv.KelvinToFahrenheit(kelv)
		} else if output == "C" {
			result = conv.KelvinToCelsius(kelv)
		} else {
			fmt.Println("Invalid output flag or value '0'. Choose either -C, -F or -K")
		}
	} else {
		fmt.Println("You must enter a valid input temperature scale (-C, -F or -K).")
	}

	fmt.Printf("%12.2f%s er %12.2f%s\n", input, getTemperatureScale(fahr, cels, kelv), result, output)

}

func getTemperatureScale(f, c, k float64) string {
	if c != 0 {
		return "°C"
	} else if f != 0 {
		return "°F"
	} else if k != 0 {
		return "°K"
	}
	return ""
}
