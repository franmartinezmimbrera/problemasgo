// fichero qsort.go
// Implementación manual del algoritmo QuickSort
package main

import (
    "fmt"
)

// particion toma el último elemento como pivote y coloca
// los menores a la izquierda y los mayores a la derecha.
func particion(arr []int, bajo, alto int) int {

    pivote := arr[alto] // Elegimos el último elemento como pivote
	i := bajo - 1       // Índice del elemento más pequeño

	for j := bajo; j < alto; j++ {
		// Si el elemento actual es menor o igual al pivote
		if arr[j] <= pivote {
			i++
			// Swap: arr[i] y arr[j]
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	// Colocamos el pivote en su posición correcta (i + 1)
	arr[i+1], arr[alto] = arr[alto], arr[i+1]
	return i + 1
}

// quicksort función principal recursiva
func quicksort(arr []int, bajo, alto int) {

    if bajo < alto {
		pi := particion(arr, bajo, alto)
		quicksort(arr, bajo, pi-1)
		quicksort(arr, pi+1, alto)
	}
}

func main() {

    datos := []int{10, 70, 8, 90, 1000, 5}
	n := len(datos)
	fmt.Print("Vector original: ")
	for _, x := range datos {
		fmt.Printf("%d ", x)
	}
	fmt.Println()
	quicksort(datos, 0, n-1)
	fmt.Print("Vector ordenado con Quicksort: ")
	for _, x := range datos {
		fmt.Printf("%d ", x)
	}
	fmt.Println()
}