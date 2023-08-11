package app

import (
	"fmt"
)

// Options muestra las opciones disponibles del menú.
func Options() {
	fmt.Printf("%s \n", "1. sumar")
	fmt.Printf("%s \n", "2. restar")
	fmt.Printf("%s \n", "3. dividir")
	fmt.Printf("%s \n", "4. multiplicar")
	fmt.Printf("%s \n", "5. salir")
	fmt.Printf("%s", "Introduce el numero de la opcion que desees: ")
}

// Menu inicia la calculadora mostrando el menú y permitiendo interacción con el usuario.
func Menu() {
	Calculadora()
}

// Calculadora es el núcleo de la calculadora, maneja las operaciones y la interacción con el usuario.
func Calculadora() {
	for {
		Options()
		var opcion int
		var a, b float64
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			a, b = readValues()
			fmt.Printf("El resultado es: %.2f \n", Sumar(a, b))
		case 2:
			a, b = readValues()
			fmt.Printf("El resultado es: %.2f \n", Restar(a, b))
		case 3:
			a, b = readValues()
			fmt.Printf("El resultado es: %.3f \n", Division(a, b))
		case 4:
			a, b = readValues()
			fmt.Printf("El resultado es: %.3f \n", Multiplicar(a, b))
		case 5:
			fmt.Println("Saliendo de la calculadora...")
			return
		default:
			fmt.Println("No existe esa opcion...")
			continue
		}
	}
}

// readValues solicita al usuario dos números y valida la entrada.
func readValues() (float64, float64) {
	var a, b float64

	for {
		// Solicitar el primer número al usuario
		fmt.Print("Introduce el primer número: ")
		_, err := fmt.Scanln(&a)
		if err != nil {
			fmt.Println("Error:", err) // Imprimir mensaje de error si hay un problema al leer el valor
			continue                   // Volver al principio del bucle para solicitar nuevamente
		}

		// Solicitar el segundo número al usuario
		fmt.Print("Introduce el segundo número: ")
		_, err = fmt.Scanln(&b)
		if err != nil {
			fmt.Println("Error:", err) // Imprimir mensaje de error si hay un problema al leer el valor
			continue                   // Volver al principio del bucle para solicitar nuevamente
		}

		break // Salir del bucle si se ingresaron números válidos
	}

	return a, b
}
