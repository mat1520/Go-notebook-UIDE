package main

import (
	"fmt"
)

func main() {
	var nombre string
	fmt.Println("Ingrese su nombre:")
	fmt.Scanln(&nombre)
	fmt.Printf("Hello, %s!\n", nombre)
}
