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
	return "Error: usuario no encontrado"
	if Contraseña == usuarios[Nombre] {
		return "Usuario iniciado sesión: " + Nombre
	}
	return "Error: contraseña incorrecta"
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
	loginMenu()
	var opcion int
	fmt.Scanln(&opcion)
	switch {
	case opcion == 1:
		fmt.Println(iniciarSesion(Nombre, Contraseña))
	case opcion == 2:
		fmt.Println(registerUser(Nombre, Contraseña))
	}

}