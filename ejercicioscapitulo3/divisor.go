// fichero divisor.go
// Este programa calcula, dados 2 números, si a es divisor de b
package main
import ("bufio";"fmt";"os";"strconv";"strings")

// esDivisor comprueba si 'a' es divisor de 'b'
func esDivisor(a, b int64) bool {
	if a == 0 {return false}
	if a < 0 {a = -a}
	if b < 0 {b = -b}
	return b%a == 0
}
func main() {
	var a, b int64;var err error
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Introduzca valor de a (el divisor): ")
	if scanner.Scan() {
		texto := strings.TrimSpace(scanner.Text())
		a, err = strconv.ParseInt(texto, 10, 64)
		if err != nil {fmt.Println("ERROR: La entrada para 'a' no es un número entero válido.");os.Exit(1)}
	}
	fmt.Print("Introduzca valor de b (el dividendo): ")
	if scanner.Scan() {
		texto := strings.TrimSpace(scanner.Text())
		b, err = strconv.ParseInt(texto, 10, 64)
		if err != nil {fmt.Println("ERROR: La entrada para 'b' no es un número entero válido.");os.Exit(1)}
	}
	if esDivisor(a, b) {
		fmt.Printf("%d ES divisor de %d\n", a, b)
	} else {
		fmt.Printf("%d NO es divisor de %d\n", a, b)
	}
}