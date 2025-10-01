package main

import (
	"fmt"
	"strconv"
	"strings"
)

func parseInput(input string) []int {
	parts := strings.Split(input, ",")
	var result []int
	for _, part := range parts {
		num, err := strconv.Atoi(strings.TrimSpace(part))
		if err == nil {
			result = append(result, num)
		}
	}
	return result
}

func sumAll(slice1 []int, slice2 []int) []int {
	if len(slice1) != len(slice2) {
		return nil
	}
	var result []int
	for i := 0; i < len(slice1); i++ {
		result = append(result, slice1[i]+slice2[i])
	}
	return result
}

func unirSlices(slice1 []int, slice2 []int) []int {
	return append(slice1, slice2...)
}

func sumOnlySlice(slice []int) int {
	sum := 0
	for _, num := range slice {
		sum += num
	}
	return sum
}

func main() {

	var slice1 []int
	var slice2 []int

	fmt.Println("Sumando dos slices")
	fmt.Println("Ingresa la lista del primer slice:")
	var input1 string
	fmt.Scanln(&input1)
	slice1 = parseInput(input1)

	fmt.Println("Ingresa la lista del segundo slice:")
	var input2 string
	fmt.Scanln(&input2)
	slice2 = parseInput(input2)

	result := sumAll(slice1, slice2)
	if result == nil {
		fmt.Println("Error: Los slices deben tener la misma longitud.")
	} else {
		fmt.Println("El resultado de la suma es:", result)
	}

	fmt.Println("Uniendo dos slices")
	combined := unirSlices(slice1, slice2)
	fmt.Println("El slice combinado es:", combined)

	fmt.Println("Sumando todos los elementos de los slices")
	sum1 := sumOnlySlice(slice1)
	sum2 := sumOnlySlice(slice2)
	fmt.Println("La suma del primer slice es:", sum1)
	fmt.Println("La suma del segundo slice es:", sum2)
	fmt.Println("La suma total de ambos slices es:", sum1+sum2)

}
