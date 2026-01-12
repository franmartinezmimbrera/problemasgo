// fichero cambiagrados.go
// Este programa cambia grados centígrados por Fahrenheit
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

  const factorCToF = 9.0 / 5.0
	fmt.Print("Introduzca valor en grados Centígrados : ")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		texto := scanner.Text()
		texto = strings.Replace(texto, ",", ".", 1)
		c, err := strconv.ParseFloat(texto, 64)
		if err != nil {
			fmt.Println("Error: El valor debe de ser un número.")
			os.Exit(1)
		}
		f := (factorCToF * c) + 32
		fmt.Printf("El resultado en grados Fahrenheit es: %v\n", f)
	}
}