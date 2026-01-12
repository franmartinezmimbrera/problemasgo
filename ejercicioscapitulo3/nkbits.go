// fichero nkbits.go
// Procedimiento que escribe los últimos k bits de un número entero.
package main
import ("fmt";"os";"strconv" )

func mostrarUltimosKBits(numero int, k int) {
	// strconv.IntSize nos da exactamente ese valor (32 o 64).
	tamanoInt := strconv.IntSize
	if k <= 0 {
		fmt.Fprintln(os.Stderr, "Error: k debe ser mayor que 0.")
		return
	}
	// Si k es mayor que los bits que tiene el número, lo limitamos
	if k > tamanoInt {
		k = tamanoInt
	}
	fmt.Printf("Los Últimos %d bits de %d son:", k, numero)

	for i := k - 1; i >= 0; i-- {
		bit := (numero >> i) & 1
		fmt.Print(bit)
	}
	fmt.Println()
}

func main() {
	num1 := 45
	mostrarUltimosKBits(num1, 8)
}