// fichero asignalibera.go
// Ejemplo de asignación dinámica de memoria usando Slices
package main
import (
	"fmt"
)
func main() {
	N := 5

	// Asignación de Memoria
	// Usamos la función incorporada 'make'.
	// make([]int, N) crea un array subyacente de 5 enteros y nos devuelve un "Slice" que apunta a él.
	// El Slice es dinámico; el tamaño no necesita ser constante en compilación.
	vector := make([]int, N)
	fmt.Printf("Vector Dinámico (Slice) de %d Elementos\n", N)
	for i := 0; i < N; i++ {
		vector[i] = i * 10
		fmt.Printf("%d ", vector[i])
	}
	fmt.Println()
	// Liberación de Memoria
	// ¡NO HACEMOS NADA!
	// Go tiene Garbage Collector.
	fmt.Println("Memoria gestionada automáticamente por el Garbage Collector (No hace falta delete).")
}