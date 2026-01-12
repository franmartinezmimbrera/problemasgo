// fichero radixsort.go
// Implementación de RadixSort (LSD - Least Significant Digit)
package main
import ("fmt")

func getMax(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}
// countingSort realiza un ordenamiento estable basado en el dígito representado por 'exp'
// exp = 1 (unidades), 10 (decenas), 100 (centenas)...
func countingSort(arr []int, exp int) {
	n := len(arr)
	output := make([]int, n) 
	count := make([]int, 10) 

	// Contar frecuencias del dígito actual
	for i := 0; i < n; i++ {
		digit := (arr[i] / exp) % 10
		count[digit]++
	}
	//  Transformar conteo en posiciones (acumulado)
	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}
	// IMPORTANTE: Recorremos el array original HACIA ATRÁS (n-1 hasta 0).
	// Esto garantiza que el algoritmo sea ESTABLE (Stable Sort).
	// Si dos números tienen el mismo dígito, el que aparecía después se mantiene después.
	for i := n - 1; i >= 0; i-- {
		digit := (arr[i] / exp) % 10
		output[count[digit]-1] = arr[i]
		count[digit]--
	}
	copy(arr, output)
}
// radixSort función principal
func radixSort(arr []int) {
	if len(arr) == 0 {
		return
	}
	m := getMax(arr)
	for exp := 1; m/exp > 0; exp *= 10 {
		countingSort(arr, exp)
	}
}
func main() {

    arr := []int{170, 45, 75, 90, 802, 24, 2, 66}
	fmt.Println("Array original:")
	fmt.Println(arr)
	radixSort(arr)
	fmt.Println("\nArray ordenado:")
	fmt.Println(arr)
}