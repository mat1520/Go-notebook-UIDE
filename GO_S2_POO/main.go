package main

import (
	"fmt"
	"math/rand"
)

func par(num int) {
	if num%2 == 0 {
		fmt.Println("El número es par")
		fmt.Println(num + rand.Intn(8))
	}
}

func impar(num int) {
	if num%2 != 0 {
		fmt.Println("El número es impar")
		fmt.Println(num - rand.Intn(80))
	}
}

func main() {
	//solicitud datos
	var n int
	fmt.Println("Valor inicial n:", n)
	fmt.Scan(&n)
	fmt.Println("Valor de n:", n)

	//funcion par
	par(n)
	//funcion impar
	impar(n)

}
