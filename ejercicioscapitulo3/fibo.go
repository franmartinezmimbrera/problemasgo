// fichero fibo.go
// Función iterativa para calcular el n-ésimo número de Fibonacci
package main
import ("fmt";"os")
// fibonacciIterativo calcula la serie usando un bucle.
func fibonacciIterativo(n int) int64 {
	var a int64 = 0
	var b int64 = 1
	var resultado int64 = 0
	if n <= 0 {return 0}
	if n == 1 {return 1}

	for i := 2; i <= n; i++ {
		resultado = a + b
		a = b
		b = resultado
	}
	return resultado
}
func main() {
	const N = 45
	resultado := fibonacciIterativo(N)
	fmt.Printf("El %d-ésimo número de Fibonacci (iterativo) es: %d\n", N, resultado)
	// El tipo int64 soporta hasta Fib(92).
	// Fib(93) ya desborda (Overflow).
	if N > 92 {
		fmt.Fprintf(os.Stderr, "Advertencia: N > 92. El resultado (%d) probablemente es incorrecto debido a desbordamiento de 'int64'.\n", resultado)
		os.Exit(1)
	}
}