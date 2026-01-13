//fichero mi_cat.go
package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// os.Args[0] es el nombre del programa, Args[1] el primer argumento
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run mi_cat.go <fichero>")
		os.Exit(1)
	}

	filename := os.Args[1]

	// Abrir el fichero
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error al abrir el fichero: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Copiar el contenido del fichero a la salida estándar (Terminal)
	// io.Copy es muy eficiente, no carga todo en memoria RAM a la vez
	_, err = io.Copy(os.Stdout, file)
	if err != nil {
		fmt.Println("Error leyendo fichero:", err)
	}
}
