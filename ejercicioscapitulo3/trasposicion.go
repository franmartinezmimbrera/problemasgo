// fichero trasposicion.go
// Función para trasponer una matriz MxN
package main
import ("fmt")

const MFilas = 3
const NColumnas = 4

type MatrizOriginal [MFilas][NColumnas]int
type MatrizTranspuesta [NColumnas][MFilas]int

func trasponerMatriz(original MatrizOriginal) MatrizTranspuesta {
	var transpuesta MatrizTranspuesta
	for i := 0; i < MFilas; i++ {
		for j := 0; j < NColumnas; j++ {
			transpuesta[j][i] = original[i][j]
		}
	}
	return transpuesta
}
func imprimirMatrizOriginal(matriz MatrizOriginal) {
	for i := 0; i < MFilas; i++ {
		for j := 0; j < NColumnas; j++ {
			fmt.Printf("%4d", matriz[i][j])
		}
		fmt.Println() // Salto de línea al terminar la fila
	}
}
func imprimirMatrizTranspuesta(matriz MatrizTranspuesta) {
	for i := 0; i < NColumnas; i++ {
		for j := 0; j < MFilas; j++ {
			fmt.Printf("%4d", matriz[i][j])
		}
		fmt.Println()
	}
}
func main() {
	A := MatrizOriginal{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
	}
	fmt.Printf("Matriz Original A (%d x %d) \n", MFilas, NColumnas)
	imprimirMatrizOriginal(A)
	AT := trasponerMatriz(A)
	fmt.Printf("\n Matriz Transpuesta AT (%d x %d) \n", NColumnas, MFilas)
	imprimirMatrizTranspuesta(AT)
}