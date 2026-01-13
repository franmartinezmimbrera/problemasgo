//fichero renombrador.go
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	// Simulación: Creamos ficheros de prueba
	os.Create("imagen1.txt")
	os.Create("imagen2.txt")
	
	patron := "*.txt"
	archivos, _ := filepath.Glob(patron)

	fmt.Println("Renombrando ficheros .txt a .bak...")

	for _, viejo := range archivos {
		// Cambiamos extensión
		nuevo := strings.TrimSuffix(viejo, ".txt") + ".bak"
		
		err := os.Rename(viejo, nuevo)
		if err == nil {
			fmt.Printf("%s -> %s\n", viejo, nuevo)
		}
	}
	
	// Limpieza para el ejercicio
	files, _ := filepath.Glob("*.bak")
	for _, f := range files { os.Remove(f) }
}
