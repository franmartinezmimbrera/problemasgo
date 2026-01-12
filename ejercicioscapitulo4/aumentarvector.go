// fichero aumentarvector.go
// Ejemplo de redimensionado dinámico usando append
package main
import (
	"fmt"
)
func main() {
    vector := make([]int, 3)
	for i := 0; i < len(vector); i++ {
		vector[i] = i + 10
	}
	fmt.Println("Vector Inicial:", vector)
	// len() es la longitud visible, cap() es la capacidad real del array subyacente
	fmt.Printf("Info Técnica -> Longitud: %d | Capacidad: %d\n", len(vector), cap(vector))
	// Redimensionado (La magia de Go)
	// Queremos llegar a 5 elementos.
	// En Go usamos 'append'.
	nInicial := 3
	nNuevo := 5
	fmt.Println("\n Ejecutando append ...")
	for i := nInicial; i < nNuevo; i++ {
		nuevoValor := i + 100
		vector = append(vector, nuevoValor)
	}
	fmt.Println("Vector Redimensionado:", vector)
	// Fíjate como la capacidad ha crecido (probablemente a 6 u 8, Go suele doblar la memoria)
	fmt.Printf("Info Técnica -> Longitud: %d | Capacidad: %d\n", len(vector), cap(vector))
    // Liberación. El Garbage Collector lo gestiona todo.
	fmt.Println("Memoria gestionada automáticamente.")
}