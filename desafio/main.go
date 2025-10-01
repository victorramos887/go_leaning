package main

import "fmt"

func main() {
	println("Digite a temperatura em Celsius:")

	var celsius float64
	fmt.Scanln(&celsius)

	var kelvin float64

	if celsius == 0 {
		kelvin = 100 + 273.15
		fmt.Println("A de ebulição em Kelvin é:", kelvin)
	} else {
		kelvin = celsius + 273.15
		fmt.Println("A temperatura em Kelvin é:", kelvin)
	}
}