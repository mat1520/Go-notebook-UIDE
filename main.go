package main

import (
	"fmt"
)

var usuarios = make(map[string]string)

func registerUser(Nombre string, Contraseña string) string {
	fmt.Println("Ingrese su nombre:")
	fmt.Scanln(&Nombre)
	fmt.Println("Ingrese su contraseña:")
	fmt.Scanln(&Contraseña)
	usuarios[Nombre] = Contraseña
	return "Usuario registrado: " + Nombre
}

func iniciarSesion(Nombre string, Contraseña string) string {
	fmt.Println("Ingrese su nombre:")
	fmt.Scanln(&Nombre)
	fmt.Println("Ingrese su contraseña:")
	fmt.Scanln(&Contraseña)
	if Nombre == usuarios[Nombre] {
		return "Usuario iniciado sesión: " + Nombre
	}
	return "Error: usuario o contraseña incorrecta"
}

func loginMenu() {
	fmt.Println("1. Iniciar sesión")
	fmt.Println("2. Registrarse")
	fmt.Print("Seleccione una opción: ")
}

func main() {

	var Nombre string
	var Contraseña string

	fmt.Println("Bienvenido, Por favor registrate o inicia sesión.")
	for {
		// Código que se ejecutará para siempre hasta que se use 'break'
		loginMenu()
		var opcion int
		fmt.Scanln(&opcion)
		switch opcion {
		case 1:
			result := iniciarSesion(Nombre, Contraseña)
			fmt.Println(result)
			if result[:24] == "Usuario iniciado sesión:" {
				fmt.Printf("Listo para transferir tu dinero, %s\n", Nombre)
				break
			}
		case 2:
			fmt.Println(registerUser(Nombre, Contraseña))
		default:
			fmt.Println("Opción no válida")
		}
	}

}
