// fichero burbuja.go
// Implementación del algoritmo de ordenación de la Burbuja (Bubble Sort)
package main
import ("fmt")
// ordenarBurbuja ordena un slice de enteros de menor a mayor.
// En Go, los Slices son "referencias" a un array subyacente.
// Al modificar los elementos dentro de la función, se modifican en el original.
func ordenarBurbuja(vec []int) {
	n := len(vec)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			
			if vec[j] > vec[j+1] {
				// INTERCAMBIO (SWAP) IDIOMÁTICO EN GO
				vec[j], vec[j+1] = vec[j+1], vec[j]
			}
		}
	}
}
func main() {
	numeros := []int{10, 7, 8, 9, 1, 5}
	fmt.Print("\nConjunto original: ")
	for _, x := range numeros {fmt.Printf("%d ", x)}
	fmt.Println()
	ordenarBurbuja(numeros)
	fmt.Print("Conjunto ordenado (Bubble Sort): ")
	for _, x := range numeros {fmt.Printf("%d ", x)}
	fmt.Println()
}