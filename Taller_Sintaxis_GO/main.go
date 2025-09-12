package main

import (
	"fmt"
	"math"
)

func preguntarRepetir() bool {
	var repetir string
	fmt.Println("¿Deseas realizar otra operación? (si/no)")
	fmt.Scan(&repetir)
	return repetir == "si" || repetir == "sí"
}

func main() {
	fmt.Println("Hecho por: Ariel Melo")

	var name string
	fmt.Println("INGRESA TU NOMBRE:")
	fmt.Scan(&name)
	fmt.Printf("BIENVENIDO %s\n", name)

	for {
		// MENU
		fmt.Printf("Hola %s, este es un programa en GO que realiza operaciones matemáticas y genera la serie de Fibonacci.\n", name)
		fmt.Println("Elige una opción:")
		fmt.Println("1. Operaciones Matematicas")
		fmt.Println("2. Serie de Fibonacci")

		var choice int
		fmt.Scan(&choice)

		// Ejecutar la opción seleccionada

		switch choice {
		case 1:
			operacionesMatematicas()
		case 2:
			serieFibonacci()
		default:
			fmt.Println("Opción no válida.")
		}

		if !preguntarRepetir() {
			fmt.Printf("Gracias por usar el programa %s.\n", name)
			break
		}
	}
}

func operacionesMatematicas() {
	fmt.Println("OPERACIONES MATEMATICAS")

	//USO DE LA LIBRERIA MATH
	//DATO FLOAT64
	numero_raiz_cuadrada := 10.20
	raiz_cuadrada := math.Sqrt(numero_raiz_cuadrada)

	fmt.Printf("La raíz cuadrada de %f es %f\n", numero_raiz_cuadrada, raiz_cuadrada)

	//DATO INT
	numero_coseno := 12
	coseno := math.Cos(float64(numero_coseno))
	fmt.Printf("La raíz cuadrada de %v es %f\n", numero_coseno, coseno)

	//DATOS INT8
	var numero_exponente int8 = 3
	base := 2.0
	exponente := float64(numero_exponente)
	potencia := math.Pow(base, exponente)
	fmt.Printf("El resultado de elevar %f a la potencia %d es %f\n", base, numero_exponente, potencia)

	//DATO FLOAT32
	numero_logaritmo := float32(100.0)
	logaritmo := math.Log(float64(numero_logaritmo))
	fmt.Printf("El logaritmo natural de %f es %f\n", numero_logaritmo, logaritmo)

	//DATO INT16
	numero_absoluto := int16(-50)
	absoluto := math.Abs(float64(numero_absoluto))
	fmt.Printf("El valor absoluto de %d es %f\n", numero_absoluto, absoluto)

	fmt.Println("Fin del PRIMER EJERCICIO")
}

func serieFibonacci() {
	fmt.Println("SERIE DE FIBONACCI")

	var n int
	fmt.Println("INGRESA UN NUMERO ENTERO DE LOS ELEMENTOS DE FIBONACCI QUE DESEAS: ")
	fmt.Scan(&n)

	a, b := 0, 1
	fmt.Printf("Serie de Fibonacci hasta %d terminos:\n", int(n))
	for i := 0; i < int(n); i++ {
		fmt.Printf("%d ", a)
		a, b = b, a+b
	}
	fmt.Println("\nFin del SEGUNDO EJERCICIO")

}
