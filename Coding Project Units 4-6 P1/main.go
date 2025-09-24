package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/fatih/color"
)

func main() {

	var entradaUsuario []string
	reader := bufio.NewReader(os.Stdin) // Create a reader to read input from the user

	color.Cyan("==========================================================")
	color.Green(" Welcome to Ariel's Units 4-6 P1 Coding Project exercise")
	color.Cyan("==========================================================\n")
	color.Red(" Recomendation: Enter two words to test how this program works")

	for {

		color.Magenta("String to add: ")

		inputString, err := reader.ReadString('\n')
		if err != nil {
			color.Red("Error reading input: %v", err)
			continue
		}

		inputString = strings.TrimSpace(inputString)
		if inputString == "" {
			color.Red("Input cannot be empty. Please enter a valid string.")
			continue
		}
		entradaUsuario = append(entradaUsuario, inputString)

		var continuar string

		for {
			color.Yellow("For continue, please confirm with? [Y/n]: ")

			continuar, err = reader.ReadString('\n')
			if err != nil {
				color.Red("Error reading input: %v", err)
				continue
			}

			continuar = strings.TrimSpace(strings.ToLower(continuar))

			if continuar == "y" || continuar == "yes" {
				break

			} else if continuar == "n" || continuar == "no" {
				goto FIN
			} else {
				color.Red("Please enter 'y', 'yes', 'n' or 'no'.")

			}
		}
	}

FIN:
	result := strings.Join(entradaUsuario, ", ")

	color.Yellow("Final result:\n")
	color.Red(result + "\n")
	color.Cyan("==========================================================")
	color.Green(" Thank you for using Ariel's coding project exercise app ")
	color.Cyan("==========================================================")
}
