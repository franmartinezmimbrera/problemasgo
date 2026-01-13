//fichero hola_flags.go
package main

import (
	"flag"
	"fmt"
)

func main() {
	// Definimos las banderas (nombre, valor_defecto, descripcion)
	nombre := flag.String("nombre", "Mundo", "Nombre a saludar")
	edad := flag.Int("edad", 0, "Edad del usuario")
	esHacker := flag.Bool("hacker", false, "Modo hacker activado")

	// Importante: Parse procesa los argumentos de la terminal
	flag.Parse()

	fmt.Printf("Hola, %s.\n", *nombre) // nombre es un puntero, usamos *
	
	if *edad > 0 {
		fmt.Printf("Tienes %d años.\n", *edad)
	}

	if *esHacker {
		fmt.Println("ACCESO AL SISTEMA CONCEDIDO... 1010101")
	}
}
// Ejecutar: go run hola_flags.go -nombre="Fran" -edad=30 -hacker
