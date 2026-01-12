// fichero binterpola.go
// Implementación del algoritmo de Búsqueda por Interpolación
package main
import ("fmt")
// busquedaInterpolacion busca un elemento estimando su posición basándose en su valor.
// Requiere que el array esté ORDENADO y uniformemente distribuido.
func busquedaInterpolacion(arr []int, objetivo int) int {
	bajo := 0
	alto := len(arr) - 1
	// El bucle continúa mientras el objetivo esté dentro del rango posible de valores
	// arr[bajo] <= objetivo <= arr[alto]
	for bajo <= alto && objetivo >= arr[bajo] && objetivo <= arr[alto] {
		// Caso base: si el rango se ha reducido a un solo punto
		if bajo == alto {
			if arr[bajo] == objetivo {
				return bajo
			}
			return -1
		}
		// Si los valores extremos son iguales, evitamos división por cero
		if arr[alto] == arr[bajo] {
			if arr[bajo] == objetivo { return bajo }
			return -1
		}
		// FÓRMULA DE INTERPOLACIÓN:
		// pos = bajo + ((alto - bajo) / (arr[alto] - arr[bajo])) * (objetivo - arr[bajo])
		// En Go, la división de enteros devuelve un entero (trunca).
		// Para obtener la precisión decimal necesaria para la interpolación,
		// DEBEMOS convertir explícitamente a float64.
		numerador := float64(alto - bajo)
		denominador := float64(arr[alto] - arr[bajo])
		distanciaValor := float64(objetivo - arr[bajo])
		fraction := numerador / denominador
		pos := bajo + int(fraction * distanciaValor)
		if pos < bajo { pos = bajo }
		if pos > alto { pos = alto }
		if arr[pos] == objetivo {
			return pos
		}
		if arr[pos] < objetivo {
			bajo = pos + 1
		} else {
			alto = pos - 1
		}
	}
	return -1
}
func main() {
	datosOrdenados := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	objetivo1 := 70
	indice := busquedaInterpolacion(datosOrdenados, objetivo1)
	fmt.Printf("\nResultado para %d: ", objetivo1)
	if indice != -1 {
		fmt.Printf("Encontrado en el índice %d.\n", indice)
	} else {
		fmt.Println("No encontrado.")
	}
}