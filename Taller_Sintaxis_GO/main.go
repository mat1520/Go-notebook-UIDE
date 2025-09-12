package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hecho por: Ariel Melo")

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

}
