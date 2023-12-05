package main

import "fmt"

func main() {
	// Get the temperature and unit from the user
	fmt.Print("Enter the temperature: ")
	var temperature float64
	fmt.Scanf("%f", &temperature)

	fmt.Print("Enter the unit. Type C, F, or K: ")
	var unit string
	fmt.Scanln(&unit)

	// Convert and display the result
	convertedTemperature, convertedUnit := convertTemperature(temperature, unit)
	fmt.Printf("Converted temperature: %.2f %s\n", convertedTemperature, convertedUnit)
}

func convertTemperature(temperature float64, unit string) (float64, string) {
	// Your code to convert temperature goes here

	if unit == "C" {
		return (temperature * 9 / 5) + 32, "Fahrenheit"
	} else if unit == "F" {
		return (temperature - 32) * 5 / 9, "Celsius"
	} else if unit == "K" {
		return temperature + 273.15, "Celsius"
	}

	// Default case if the unit is not recognized
	fmt.Println("Invalid unit. Please enter C, F, or K.")
	return temperature, unit
}
