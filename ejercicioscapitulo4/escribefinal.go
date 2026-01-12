// fichero escribefinal.go
// Ejemplo para escribir al final de un fichero de texto (Append)
package main
import (
	"fmt"
	"os"
)
func main() {
	// Abrir fichero con banderas específicas (Flags)
	// os.O_APPEND: Escribir al final sin borrar lo anterior.
	// os.O_WRONLY: Solo escritura.
	// os.O_CREATE: Si no existe, créalo (comportamiento robusto).
	// 0644: Permisos del archivo (Lectura/Escritura para mí, Lectura para el resto).
	fichero, err := os.OpenFile("datos.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error al abrir el fichero datos.txt en modo append: %v\n", err)
		os.Exit(1)
	}
	defer fichero.Close()
	// Escribir al final
	_, err = fichero.WriteString("Esta es la línea AÑADIDIDA al final (Append)\n")
	
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error al escribir en el fichero: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("\nSe ha añadido una línea al final de 'datos.txt'.")
}