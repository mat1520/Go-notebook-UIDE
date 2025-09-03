package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nombre string
	var edad int
	fmt.Println("Ingrese su nombre:")
	fmt.Scanln(&nombre)
	fmt.Println("Ingrese su edad:")
	fmt.Scanln(&edad)
	fmt.Printf("Hello, %s! You are %d years old.\n", nombre, edad)

	var myInt int = 12
	fmt.Println(myInt)
	fmt.Println(reflect.TypeOf(float64(myInt)))
}
