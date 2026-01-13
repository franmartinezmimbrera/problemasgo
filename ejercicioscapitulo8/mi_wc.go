//fichero mi_wc.go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Por favor, indica un fichero.")
		return
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineas := 0
	palabras := 0
	caracteres := 0
	for scanner.Scan() {
		linea := scanner.Text()
		lineas++
		caracteres += len(linea)
		// Fields divide por espacios en blanco
		palabras += len(strings.Fields(linea))
	}

	fmt.Printf("Líneas: %d\nPalabras: %d\nCaracteres (aprox): %d\n", lineas, palabras, caracteres)
}
