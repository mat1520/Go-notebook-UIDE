package main

import (
	"fmt"
)

func main() {
	var nombre string
	var edad int
	fmt.Println("Ingrese su nombre:")
	fmt.Scanln(&nombre)
	fmt.Println("Ingrese su edad:")
	fmt.Scanln(&edad)
	fmt.Printf("Hello, %s! You are %d years old.\n", nombre, edad)
}
