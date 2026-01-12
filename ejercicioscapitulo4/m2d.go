// fichero m2d.go
// Ejemplo de matrices dinámicas (Slice de Slices)
package main

import (
	"fmt"
)
func main() {
	M := 3 // Filas
	N := 4 // Columnas

	matriz := make([][]int, M)
	for i := 0; i < M; i++ {
		// Asignamos memoria para la fila 'i' con N columnas
		matriz[i] = make([]int, N)
		for j := 0; j < N; j++ {
			matriz[i][j] = i*N + j + 1
		}
	}
	fmt.Printf("\nMatriz Dinámica %dx%d\n", M, N)
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%3d ", matriz[i][j])
		}
		fmt.Println()
	}
	// Liberación
	// ¡NADA!
	// El Garbage Collector detecta cuando 'matriz' deja de usarse.
	fmt.Println("Memoria gestionada automáticamente por Go.")
}