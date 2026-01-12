// fichero raizcuadrada.go
// Este programa calcula la raíz cuadrada de un número
package main

import (
	"bufio"
	"fmt"
	"math" 
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Introduce el número a calcular la raíz cuadrada: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		texto := scanner.Text()
		texto = strings.Replace(texto, ",", ".", 1)
		numero, err := strconv.ParseFloat(texto, 64)
		if err != nil {
			fmt.Println("Error: El valor debe de ser un número.")
			os.Exit(1)
		}
		if numero < 0 {
			fmt.Println("Error: El valor debe ser un número positivo")
			os.Exit(1)
		}
		resultado := math.Sqrt(numero)
		fmt.Printf("La raíz cuadrada de %v es: %v\n", numero, resultado)
	}
}