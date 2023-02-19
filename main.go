package main

import (
	"flag"
	"fmt"

	"conv_test.go/conv"
)

var (
	fahr    float64
	celsius float64
	kelvin  float64
	out     string
)

func init() {
	flag.Float64Var(&fahr, "F", 0.0, "temperature in degrees Fahrenheit")
	flag.Float64Var(&celsius, "C", 0.0, "temperature in degrees Celsius")
	flag.Float64Var(&kelvin, "K", 0.0, "temperature in degrees Kelvin")
	flag.StringVar(&out, "out", "C", "output unit (C - Celsius, F - Fahrenheit, K - Kelvin)")
}

func main() {
	flag.Parse()

	switch out {
	case "C":
		if isFlagPassed("F") {
			celsius = conv.FahrenheitToCelsius(fahr)
		} else if isFlagPassed("K") {
			celsius = conv.KelvinToCelsius(kelvin)
		}
		fmt.Printf("%.2f degrees %s is %.2f degrees Celsius\n", getTemperature(), getUnit(), celsius)
	case "F":
		if isFlagPassed("C") {
			fahr = conv.CelsiusToFahrenheit(celsius)
		} else if isFlagPassed("K") {
			fahr = conv.KelvinToFahrenheit(kelvin)
		}
		fmt.Printf("%.2f degrees %s is %.2f degrees Fahrenheit\n", getTemperature(), getUnit(), fahr)
	case "K":
		if isFlagPassed("C") {
			kelvin = conv.CelsiusToKelvin(celsius)
		} else if isFlagPassed("F") {
			kelvin = conv.FahrenheitToKelvin(fahr)
		}
		fmt.Printf("%.2f degrees %s is %.2f Kelvin\n", getTemperature(), getUnit(), kelvin)
	default:
		fmt.Println("Invalid output unit specified")
	}
}

func getTemperature() float64 {
	switch {
	case isFlagPassed("F"):
		return fahr
	case isFlagPassed("C"):
		return celsius
	case isFlagPassed("K"):
		return kelvin
	default:
		return 0.0
	}
}

func getUnit() string {
	switch {
	case isFlagPassed("F"):
		return "Fahrenheit"
	case isFlagPassed("C"):
		return "Celsius"
	case isFlagPassed("K"):
		return "Kelvin"
	default:
		return ""
	}
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
