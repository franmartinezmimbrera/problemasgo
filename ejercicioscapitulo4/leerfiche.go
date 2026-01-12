// fichero leerfiche.go
// Ejemplo para leer ficheros de texto línea a línea
package main
import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// os.Open abre el archivo en modo "solo lectura" (Read Only).
	fichero, err := os.Open("datos.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error al abrir el fichero datos.txt: %v\n", err)
		os.Exit(1)
	}
	defer fichero.Close()

	fmt.Println("\nContenido de 'datos.txt':")
	// El Scanner es un "iterador" que lee datos (por defecto, línea por línea)
	// Actúa como buffer intermedio para no hacer demasiadas llamadas al sistema.
	scanner := bufio.NewScanner(fichero)
	for scanner.Scan() {
		linea := scanner.Text()
		fmt.Println(linea)
	}
	// Siempre hay que comprobar scanner.Err() después del bucle.
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error leyendo el fichero: %v\n", err)
	}
}