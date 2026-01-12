// fichero bbinaria.go
// Implementación del algoritmo de Búsqueda Binaria (Iterativa)
package main

import ("fmt";"slices" )

// busquedaBinaria asume que el slice 'arr' está ORDENADO de menor a mayor.
// Devuelve el índice del elemento o -1 si no existe.
func busquedaBinaria(arr []int, objetivo int) int {
	bajo := 0
	alto := len(arr) - 1
	for bajo <= alto {
		// Calculamos el punto medio.
		// Usamos "bajo + (alto-bajo)/2" en lugar de "(bajo + alto)/2"		
		medio := bajo + (alto-bajo)/2
		if arr[medio] == objetivo {
			return medio // ¡Encontrado!
		}
		if arr[medio] < objetivo {
			// El objetivo está en la mitad superior
			bajo = medio + 1
		} else {
			// El objetivo está en la mitad inferior
			alto = medio - 1
		}
	}
	return -1 // No encontrado
}

func main() {
	
	datosOrdenados := []int{5, 8, 12, 15, 20, 30, 40}
	objetivo1 := 20
	fmt.Printf("Conjunto de datos (Ordenado): %v\n", datosOrdenados)
	indice := busquedaBinaria(datosOrdenados, objetivo1)
	if indice != -1 {
		fmt.Printf("Resultado para %d: Encontrado (función manual) en el índice %d.\n", objetivo1, indice)
	} else {
		fmt.Printf("Resultado para %d: No encontrado.\n", objetivo1)
	}
	// slices.BinarySearch devuelve el índice y un booleano 'found'
	idx, found := slices.BinarySearch(datosOrdenados, objetivo1)
	if found { indice = idx } else { indice = -1 }
	if indice != -1 {
		fmt.Printf("Resultado para %d: Encontrado (función de slices) en el índice %d.\n", objetivo1, indice)
	} else {
		fmt.Printf("Resultado para %d: No encontrado.\n", objetivo1)
	}
}