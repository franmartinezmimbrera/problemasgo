// fichero factoriali2.go
// Calcula el factorial de un número entero no negativo de forma iterativa
package main
import ("fmt";"math/big")

func main() {

	resultado := big.NewInt(1)
	limite := 100
	
	for i := 1; i <= limite; i++ {		
		iBig := big.NewInt(int64(i))
		resultado.Mul(resultado, iBig)
	}
	
	fmt.Printf("El factorial de %d es:\n%s\n", limite, resultado.String())
}