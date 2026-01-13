//fichero variables_entorno.go
package main

import (
	"fmt"
	"os"
)

func main() {
	// Obtener una variable de entorno específica
	usuario := os.Getenv("USER") // En Windows suele ser USERNAME
	if usuario == "" {
		usuario = os.Getenv("USERNAME")
	}

	path := os.Getenv("PATH")

	fmt.Printf("Usuario del sistema: %s\n", usuario)
	fmt.Printf("Primeras rutas del PATH: %.50s...\n", path)

	// Establecer una variable temporalmente para este proceso
	os.Setenv("MI_APP_PORT", "8080")
	fmt.Println("Puerto configurado:", os.Getenv("MI_APP_PORT"))
}
