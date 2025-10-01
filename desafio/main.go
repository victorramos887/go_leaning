package main

import (
	"fmt"
	"strconv"
)

func main() {
	println("Digite a temperatura em Celsius:")

	var celsius string
	fmt.Scanln(&celsius)

	var kelvin float64

	if celsius == "" {
		kelvin = 100 + 273.15
		fmt.Println("A temperatura de ebulição em Kelvin é:", kelvin)
	} else {
		temp, err := strconv.ParseFloat(celsius, 64)
		if err != nil {
			fmt.Println("Erro: valor inválido. Usando temperatura de ebulição.")
			kelvin = 100 + 273.15
			fmt.Println("A temperatura de ebulição em Kelvin é:", kelvin)
		} else {
			kelvin = temp + 273.15
			fmt.Printf("A temperatura %.2f°C em Kelvin é: %.2f K\n", temp, kelvin)
		}
	}
}