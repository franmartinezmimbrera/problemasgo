// fichero multiplicamatrices.go
// Multiplicación de dos matrices cuadradas de tamaño fijo.
package main
import ("fmt")

const N = 3
type Matriz [N][N]int

// Para matrices pequeñas (3x3) está bien, para grandes usaríamos punteros *Matriz.
func multiplicarMatrices(a, b Matriz) Matriz {
	var c Matriz 
	for i := 0; i < N; i++ {          
		for j := 0; j < N; j++ {      
			for k := 0; k < N; k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return c
}
func imprimirMatriz(matriz Matriz) {
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			fmt.Printf("%4d ", matriz[i][j])
		}
		fmt.Println() 
	}
}
func main() {

	a := Matriz{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	b := Matriz{
		{9, 8, 7},
		{6, 5, 4},
		{3, 2, 1},
	}
	c := multiplicarMatrices(a, b)
	fmt.Printf("Matriz A (%dx%d):\n", N, N)
	imprimirMatriz(a)
	fmt.Printf("\nMatriz B (%dx%d):\n", N, N)
	imprimirMatriz(b)
	fmt.Println("\nMatriz Resultado C = A * B:")
	imprimirMatriz(c)
}