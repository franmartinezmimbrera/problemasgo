// fichero mascara.go
// Invierte (cambia 0->1 y 1->0) los últimos k bits de un número entero
package main
import ("fmt";"os")
func invertirUltimosKBits(numero int64, k int) int64 {
	const tamanoLL = 64
	if k <= 0 {return numero}
	if k >= tamanoLL {return ^numero}
	mascara := (uint64(1) << k) - 1
	return numero ^ int64(mascara)
}
func main() {
	var num int64 = 45
	k := 4
	resultado := invertirUltimosKBits(num, k)
	fmt.Printf("Número original:   %d\t(Binario: ...%08b)\n", num, num)
	fmt.Printf("K: %d\n", k)
	fmt.Printf("Número resultante: %d\t(Binario: ...%08b)\n", resultado, resultado)

}