//fichero compresor.go
package main
import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)
func main() {
	archivos := []string{"hola_flags.go", "mi_cat.go"}
	zipName := "codigos.zip"
	nuevoZip, err := os.Create(zipName)
	if err != nil { panic(err) }
	defer nuevoZip.Close()
	writer := zip.NewWriter(nuevoZip)
	defer writer.Close()

	for _, archivo := range archivos {
		fmt.Printf("Comprimiendo %s...\n", archivo)
		
		// Abrir fichero original
		f, err := os.Open(archivo)
		if err != nil {
			fmt.Println("Saltando (no existe):", archivo)
			continue
		}
		defer f.Close()

		// Crear entrada dentro del zip
		w, _ := writer.Create(archivo)
		
		// Copiar contenido
		io.Copy(w, f)
	}
	fmt.Println("Backup creado:", zipName)
}
