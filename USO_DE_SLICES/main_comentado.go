/*package main

import (
	"fmt"
	"strings"
)

func filtrarCompletadas(tareas []string) []string {
	var tareasPendientes []string
	for i := range tareas {
		if !strings.HasSuffix(tareas[i], "(completada)") {
			tareasPendientes = append(tareasPendientes, tareas[i])
		}
	}
	return tareasPendientes
}

func main() {
	tareas := []string{
		"Comprar leche",
		"Pagar la luz (completada)",
		"Llamar al dentista",
		"Enviar correo a Juan (completada)",
		"Terminar este ejercicio",
	}

	fmt.Println("Lista Original:", tareas)

	tareasPendientes := filtrarCompletadas(tareas)

	fmt.Println("Las tareas pendientes son:", tareasPendientes)

	//separar del slices para imprimir mejor
	for i, tarea := range tareasPendientes {
		fmt.Printf("%d. %s\n", i+1, tarea)
	}

}*/