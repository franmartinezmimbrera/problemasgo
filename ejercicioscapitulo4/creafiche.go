// fichero creafiche.go
// Ejemplo para crear/escribir ficheros
package main
import (
    "fmt"
    "os" 
)  

func main() {

    fichero, err := os.Create("datos.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error al abrir/crear el fichero datos.txt: %v\n", err)
		os.Exit(1)
	}
	// 'defer' asegura que fichero.Close() se ejecute cuando la función main termine,
	// pase lo que pase (incluso si hay errores intermedios).
	defer fichero.Close()

    // Escritura de línea
	fichero.WriteString("Esta es la primera línea.\n")
	pi := 3.14159
    // Escritura con formato
	fmt.Fprintf(fichero, "El número PI es aproximadamente %.4f\n", pi)
	fichero.WriteString("Tercera línea de ejemplo.\n")
	fmt.Println("El archivo 'datos.txt' fue creado y escrito correctamente.")

}