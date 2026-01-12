// fichero shellsort.go
// Implementación de ShellSort usando la secuencia de Knuth (h = 3*h + 1)
package main

import (
    "fmt"
)

// shellsortKnuth ordena el slice in-place usando saltos variables.
func shellsortKnuth(arr []int) {

    n := len(arr)
	h := 1
	// Calcular el salto inicial (Secuencia de Knuth)
	// La secuencia es: 1, 4, 13, 40, 121...
	// Buscamos el h más grande posible que sea menor que n/3
	for h < n/3 {
		h = 3*h + 1
	}
	// vamos reduciendo el salto 'h'
	for h >= 1 {
		// Insertion Sort modificado con salto 'h'
		// En lugar de comparar con el vecino inmediato (i-1), comparamos con (i-h)
		for i := h; i < n; i++ {
			temp := arr[i]
			j := i
			// Desplazamos los elementos mayores hacia la derecha (saltando h posiciones)
			// j >= h verifica que no nos salgamos del array por la izquierda
			for j >= h && arr[j-h] > temp {
				arr[j] = arr[j-h]
				j -= h
			}	
			// Colocamos el elemento temporal en su hueco
			arr[j] = temp
		}
		// Reducimos el salto según la inversa de Knuth
		h = h / 3
	}
}

func main() {

    datos := []int{90, 8, 70, 6, 50, 4, 30, 2, 10, 0, 85, 15, 65, 35}

	fmt.Println("Array original desordenado:")
	fmt.Println(datos) 
	
    shellsortKnuth(datos)
	
    fmt.Println("\nArray ordenado con Shellsort (Knuth):")
	fmt.Println(datos)

}