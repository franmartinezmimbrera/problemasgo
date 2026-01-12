// fichero qsort2.go
// Ordenación utilizando Pattern-Defeating Quicksort (pdqsort)
package main

import (
	"fmt"
	"slices" // Paquete moderno para operaciones con slices (Go 1.21+)
)

func main() {
	
	numeros := []int{50, 10, 8, 20, 40}

	fmt.Print("Vector antes de ordenar: ")
	fmt.Println(numeros)

	// Go utiliza internamente 'pdqsort' (Pattern-Defeating Quicksort),
	// un algoritmo híbrido extremadamente rápido (O(n log n)).
	
    slices.Sort(numeros)
	
    fmt.Print("Vector después de ordenar: ")
	fmt.Println(numeros)
}