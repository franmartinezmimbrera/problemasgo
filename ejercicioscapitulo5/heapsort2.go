// fichero heapsort2.go
// Ejemplo de uso de container/heap para ordenar
package main
import ("container/heap";"fmt")
// Definimos un tipo personalizado (un slice de enteros)
type IntHeap []int
// Implementamos la interfaz sort.Interface
func (h IntHeap) Len() int { return len(h) }
// Less define el orden.
// Si usamos < (menor que), construimos un Min-Heap (el menor arriba).
// Al extraer elementos de un Min-Heap uno a uno, obtendremos una lista ordenada ascendente.
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }
// Implementamos Push y Pop (requisitos de heap.Interface)
// Usa un puntero (*IntHeap) porque Push y Pop modifican el tamaño del slice.
func (h *IntHeap) Push(x any) {*h = append(*h, x.(int))}
func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]     // Tomamos el último elemento
	*h = old[0 : n-1] // Reducimos el slice
	return x
}
func main() {
	// Creamos el slice y lo convertimos en un puntero a nuestra estructura
	nums := &IntHeap{12, 11, 13, 5, 6, 7}
	fmt.Println("Datos originales:", *nums)
	// heap.Init reordena el slice para cumplir la propiedad de montículo. O(n)
	heap.Init(nums)
	fmt.Println("Tras heap.Init (Min-Heap):", *nums)
	fmt.Print("Extrayendo ordenadamente: ")
	// Extracción ordenada (HeapSort)
	// Mientras queden elementos, sacamos el mínimo (Pop)
	for nums.Len() > 0 {
		val := heap.Pop(nums)
		fmt.Printf("%d ", val)
	}
	fmt.Println()
}