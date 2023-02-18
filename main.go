package main

import (
	"danieldg91/funtemps/conv"
	"flag"
	"fmt"
	"log"
)

var (
	fahr float64
	cels float64
	kelv float64
	in   string
	out  string
)

func init() {
	flag.Float64Var(&fahr, "F", 0, "Enter a temperature in Fahrenheit")
	flag.Float64Var(&cels, "C", 0, "Enter a temperature in Celsius")
	flag.Float64Var(&kelv, "K", 0, "Enter a temperature in Kelvin")
	flag.StringVar(&in, "in", "", "Enter the input temperature unit (F, C, K)")
	flag.StringVar(&out, "out", "", "Enter the desired output temperature unit (F, C, K)")

}

func main() {

	flag.Parse()

	var input float64
	var output float64

	// Convert temperature
	if in == "C" {
		input = cels
		switch out {
		case "K":
			output = conv.CelsiusToKelvin(cels)
			break
		case "F":
			output = conv.CelsiusToFahrenheit(cels)
			break
		default:
			log.Fatalf("Conversion from Celsius to invalid flag. '-out K' or '-out F' are valid outputs.")
			break
		}
	} else if in == "F" {
		input = fahr
		switch out {
		case "C":
			output = conv.FahrenheitToCelsius(fahr)
			break
		case "K":
			output = conv.FahrenheitToKelvin(fahr)
			break
		default:
			log.Fatalf("Conversion from Fahrenheit to invalid flag. '-out C' or '-out K' are valid outputs.")
			break
		}
	} else if in == "K" {
		input = kelv
		switch out {
		case "C":
			output = conv.KelvinToCelsius(kelv)
			break
		case "F":
			output = conv.KelvinToFahrenheit(kelv)
			break
		default:
			log.Fatalf("Conversion from Celsius to invalid flag. '-out K' or '-out F' are valid outputs.")
			break
		}
	} else {
		log.Fatalf("Invalid conversion. One of the following scales must be different from your input: -C (Celsius) -K (Kelvin) -F (Fahrenheit)")
	}

	// Print results
	fmt.Printf("%12.2f°%s is %12.2f°%s\n", input, in, output, out)
}
