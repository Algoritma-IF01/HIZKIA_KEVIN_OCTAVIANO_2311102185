package main

import "fmt"

func main() {
	var celsius int
	fmt.Print("Temperatur Celsius: ")
	fmt.Scanln(&celsius)

	reamur := celsius * 4 / 5
	fahrenheit := (celsius * 9 / 5) + 32
	kelvin := celsius + 273
	
	fmt.Printf("Derajat Reamur: %d\n", reamur)
	fmt.Printf("Derajat Fahrenheit: %d\n", fahrenheit)
	fmt.Printf("Derajat Kelvin: %d\n", kelvin)
}