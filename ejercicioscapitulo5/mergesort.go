// fichero mergesort.go
// Implementación del algoritmo MergeSort (Divide y Vencerás)
package main

import ("fmt")
// funcion merge: Fusiona dos subconjuntos ordenados en un solo conjunto ordenado.
// Corresponde a la fase de "Vencer" o "Conquistar".
func merge(arr []int, inicio, medio, fin int) {

	tamano := fin - inicio + 1
	temp := make([]int, tamano)
	i := inicio    // Puntero para la mitad izquierda
	j := medio + 1 // Puntero para la mitad derecha
	k := 0         // Puntero para el array temporal
	for i <= medio && j <= fin {
		if arr[i] <= arr[j] {
			temp[k] = arr[i]
			i++
		} else {
			temp[k] = arr[j]
			j++
		}
		k++
	}
	// Si sobran elementos en la mitad izquierda, los copiamos
	for i <= medio {
		temp[k] = arr[i]
		k++
		i++
	}
	// Si sobran elementos en la mitad derecha, los copiamos
	for j <= fin {
		temp[k] = arr[j]
		k++
		j++
	}
	// Copiamos el array temporal ordenado de vuelta al array original
	// copy(destino, origen) es la forma eficiente en Go
	copy(arr[inicio:fin+1], temp)
}
// mergeSort: Función recursiva que divide el array
func mergeSort(arr []int, inicio, fin int) {
	if inicio < fin {
		// Encontramos el punto medio para dividir
		// Evitamos desbordamiento (overflow) usando esta fórmula en lugar de (inicio+fin)/2
		medio := inicio + (fin-inicio)/2
        // DIVIDIR
		mergeSort(arr, inicio, medio)
    	mergeSort(arr, medio+1, fin)
        //VENCER (Mezclar las mitades ordenadas)
		merge(arr, inicio, medio, fin)
	}
}
func main() {

	datos := []int{38, 27, 43, 3, 9, 82, 10}
	n := len(datos)
	fmt.Println("Vector original desordenado:")
	for _, x := range datos {
		fmt.Printf("%d ", x)
	}
	fmt.Println()
	mergeSort(datos, 0, n-1)
	fmt.Println("\nVector ordenado con MergeSort:")
	for _, x := range datos {
		fmt.Printf("%d ", x)
	}
	fmt.Println()
}