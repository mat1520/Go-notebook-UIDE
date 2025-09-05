package main

import "fmt"

func main() {
	// Declaramos una variable de tipo entero
	x := 10

	// Declaramos un puntero 'p' y le asignamos la dirección de memoria de 'x'
	p := &x

	// Imprimimos el valor de 'x'
	fmt.Println("Valor de x:", x) // Salida: Valor de x: 10

	// Imprimimos la dirección de memoria de 'x' usando el puntero 'p'
	fmt.Println("Dirección de memoria de x:", p) // Salida: Dirección de memoria de x: 0xc0000140a0 (el valor puede variar)

	// Imprimimos el valor de la variable a la que apunta 'p' (desreferenciación)
	fmt.Println("Valor al que apunta p:", *p) // Salida: Valor al que apunta p: 10

	// Modificamos el valor de 'x' a través del puntero 'p'
	*p = 20

	// Imprimimos el nuevo valor de 'x'
	fmt.Println("Nuevo valor de x:", x) // Salida: Nuevo valor de x: 20
}
