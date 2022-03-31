package main

import "fmt"

func convertKelvin(temperatureInput float64) (float64, float64) {
	resultF := (9 * temperatureInput / 5) - 459.67
	resultC := 5 * (resultF - 32) / 9

	return resultF, resultC
}

func convertCelsius(temperatureInput float64) (float64, float64) {
	resultF := (9*temperatureInput + 160) / 5
	resultK := 5 * (resultF + 459.67) / 9
	return resultF, resultK
}

func convertFahrenheit(temperatureInput float64) (float64, float64) {
	resultC := 5 * (temperatureInput - 32) / 9
	resultK := 5 * (temperatureInput + 459.67) / 9

	return resultC, resultK
}

func main() {
	var temperatureChoice int
	var temperatureInput float64
	fmt.Println("Enter your temperature format (1 for Kelvin, 2 for Celsius, 3 for Fahrenheit: ")
	fmt.Scanln(&temperatureChoice)
	fmt.Println("Enter the temperature: ")
	fmt.Scanln(&temperatureInput)

	if temperatureChoice == 1 {
		resultF, resultC := convertKelvin(temperatureInput)
		fmt.Println("Fahrenheit: ", resultF, " Celsius: ", resultC)
	} else if temperatureChoice == 2 {
		resultF, resultK := convertCelsius(temperatureInput)

		fmt.Println("Fahrenheit: ", resultF, " Kelvin: ", resultK)
	} else if temperatureChoice == 3 {
		resultC, resultK := convertFahrenheit(temperatureInput)

		fmt.Println("Kelvin: ", resultK, " Celsius: ", resultC)
	} else {
		fmt.Println("No Conversion")
	}

}
