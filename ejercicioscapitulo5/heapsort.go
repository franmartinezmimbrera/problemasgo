// fichero heapsort.go
// Implementación manual de HeapSort (Ordenación por Montículo)
package main

import ("fmt")

// heapify asegura que el subárbol con raíz en el índice 'i' cumpla la propiedad de Max-Heap.
// n es el tamaño del heap "activo" (va disminuyendo conforme ordenamos).
func heapify(arr []int, n int, i int) {
	largest := i       // Inicializamos el más grande como la raíz
	left := 2*i + 1    // Hijo izquierdo
	right := 2*i + 2   // Hijo derecho

	// Si el hijo izquierdo es más grande que la raíz
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	// Si el hijo derecho es más grande que el más grande hasta ahora
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	// Si el más grande no es la raíz
	if largest != i {
		// Hacemos SWAP: la raíz baja y el hijo grande sube
		arr[i], arr[largest] = arr[largest], arr[i]
		// Recursivamente aplicamos heapify al subárbol afectado
		heapify(arr, n, largest)
	}
}

// heapSort es la función principal
func heapSort(arr []int) {
	n := len(arr)

	// Construir el Heap (Build Max-Heap)
	// Empezamos desde el último nodo que NO es hoja (n/2 - 1) hacia arriba.
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}
	// Extraer elementos del heap uno a uno
	for i := n - 1; i > 0; i-- {
		// Movemos la raíz actual (el máximo) al final del array
		arr[0], arr[i] = arr[i], arr[0]
		// Llamamos a heapify en el heap reducido (tamaño i)
		// para restaurar la propiedad de Max-Heap en la raíz.
		heapify(arr, i, 0)
	}
}

func main() {

	arr := []int{12, 11, 13, 5, 6, 7}
	fmt.Println("Array original:")
	fmt.Println(arr)
	heapSort(arr)
	fmt.Println("Array ordenado con Heapsort:")
	fmt.Println(arr)
}