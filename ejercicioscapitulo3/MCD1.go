// fichero mcd1.go
// Este programa calcula el MCD dados 2 números mediante una función
package main
import ("bufio";"fmt";"os";"strconv";"strings")
// MCD calcula el Máximo Común Divisor usando el algoritmo de Euclides
func MCD(a, b int64) int64 {
	var temporal int64
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	// Algoritmo de Euclides
	for a > 0 {
		temporal = a
		a = b % a
		b = temporal
	}
	return b
}
func main() {
	var a, b int64;var err error
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\n   Introduzca valor de a: ")
	if scanner.Scan() {
		texto := strings.TrimSpace(scanner.Text())
		a, err = strconv.ParseInt(texto, 10, 64)
		if err != nil {
			fmt.Println("ERROR: La entrada no es un número entero")
			os.Exit(1)
		}
	}
	fmt.Print("\n  Introduzca valor de b: ")
	if scanner.Scan() {
		texto := strings.TrimSpace(scanner.Text())
		b, err = strconv.ParseInt(texto, 10, 64)
		if err != nil {
			fmt.Println("ERROR: La entrada no es un número entero")
			os.Exit(1)
		}
	}
	if a < 0 {a = -a}
	if b < 0 {b = -b}
	fmt.Printf("EL M.C.D de %d y %d es: %d\n", a, b, MCD(a, b))
}