// fichero factorialr.go
// Calcula el factorial de un número entero no negativo de forma recursiva.
package main
import ("bufio";"fmt";"os";"strconv";"strings")

func factorialRecursivo(n int64) int64 {
	if n <= 1 {return 1}
	return n * factorialRecursivo(n-1)
}
func main() {
	var numero int64;var err error
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Introduce un número entero no negativo (máx 20): ")
	if scanner.Scan() {
		texto := strings.TrimSpace(scanner.Text())
		numero, err = strconv.ParseInt(texto, 10, 64)
		if err != nil {fmt.Fprintln(os.Stderr, "Error: Entrada no válida.");os.Exit(1)}
	} else {
		return
	}
	if numero < 0 {
		fmt.Fprintln(os.Stderr, "Error: El factorial solo está definido para números >0")
		os.Exit(1)
	}
	if numero > 20 {
		fmt.Printf("Advertencia: El factorial de %d es muy grande y probablemente cause un resultado incorrecto por desbordamiento (overflow).\n", numero)
	}
	resultado := factorialRecursivo(numero)
	fmt.Printf("El factorial de %d es: %d\n", numero, resultado)
}