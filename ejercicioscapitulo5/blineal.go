// fichero blineal.go
// Implementación del algoritmo de Búsqueda Lineal (Secuencial)
package main
import ("fmt";"slices")

// busquedaLineal recorre el slice elemento a elemento.
// Retorna el índice si lo encuentra, o -1 si no existe.
func busquedaLineal(arr []int, objetivo int) int {
	// range devuelve (índice, valor)
	for i, v := range arr {
		if v == objetivo {
			return i // Encontrado
		}
	}
	return -1 // No encontrado tras recorrer todo
}
func main() {
	datos := []int{10, 5, 20, 15, 8, 30}
	objetivo1 := 15
	fmt.Printf("Conjunto de datos: %v\n", datos)
    //Método manual
	indice := busquedaLineal(datos, objetivo1)
    if indice != -1 {
		fmt.Printf("Resultado para %d: Encontrado en índice %d.\n", objetivo1, indice)
	} else {
		fmt.Printf("Resultado para %d: No encontrado\n", objetivo1)
	}	
    //Usando la fórmula de slices.
	indice = slices.Index(datos, objetivo1)
	if indice != -1 {
		fmt.Printf("Resultado para %d: Encontrado en índice %d.\n", objetivo1, indice)
	} else {
		fmt.Printf("Resultado para %d: No encontrado\n", objetivo1)
	}
}