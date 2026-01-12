// fichero pesetaseuros.go
// Este programa realiza la conversión de pesetas a euros
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	const tasaConversion = 166.386

	fmt.Print("Introduzca su valor en pesetas: ")
	scanner := bufio.NewScanner(os.Stdin)
	
	if scanner.Scan() {
		textoEntrada := scanner.Text()
		textoEntrada = strings.Replace(textoEntrada, ",", ".", 1)
		pesetas, err := strconv.ParseFloat(textoEntrada, 64)
		if err != nil {
			fmt.Println("ERROR: La entrada no es un número válido.")
			os.Exit(1) 
		}
		if pesetas < 0 {
			fmt.Println("ERROR: El valor en pesetas no puede ser negativo.")
			os.Exit(1)
		}

		euros := pesetas / tasaConversion
		fmt.Printf("Su valor de %.2f pesetas equivale a: %.2f euros\n", pesetas, euros)
	}
}