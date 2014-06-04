package main

import "fmt"

func main() {
	fmt.Print("Enter a number")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	fmt.Println(output)

	convertTemp()
}

func convertTemp() {
	fmt.Print("enter the temperature in deg far.: ")
	var far float64
	fmt.Scanf("%f", &far)

	cel := (far - 32) * 5/9

	fmt.Println(cel)
}