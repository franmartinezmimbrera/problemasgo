// fichero mcd.go
// Este programa calcula el MCD dados 2 números enteros usando el Algoritmo de Euclides
package main
import ("bufio";"fmt";"os";"strconv";"strings")
func main() {
	var a, b, a2, b2, temporal int64;var err error
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
	a2 = a
	b2 = b
	if a < 0 {a = -a}
	if b < 0 {b = -b}
	if a == 0 || b == 0 {
		resultado := a
		if a == 0 {
			resultado = b
		}
		fmt.Printf("\n EL M.C.D de %d y %d es: %d\n", a2, b2, resultado)
		return
	}
	// ALGORITMO DE EUCLIDES	
	for a != 0 {
		temporal = a
		a = b % a
		b = temporal
	}
	fmt.Printf(" EL M.C.D de %d y %d es: %d\n", a2, b2, b)
}