// fichero factoriali.go
// Calcula el factorial de un número entero no negativo de forma iterativa
package main
import ("bufio";"fmt";"os";"strconv";"strings")
func factorialIterativo(n int64) int64 {
	if n == 0 || n == 1 {return 1}
	resultado := int64(1)
	for i := int64(2); i <= n; i++ {resultado *= i }
	return resultado
}
func main() {
	var numero int64;var err error
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Introduce un numero entero no negativo (max 20 para precision): ")
	if scanner.Scan() {
		texto := strings.TrimSpace(scanner.Text())
		numero, err = strconv.ParseInt(texto, 10, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: Entrada no valida.")
			os.Exit(1)
		}
	} else {
		return 
	}
	if numero < 0 {
		fmt.Fprintln(os.Stderr, "Error: El factorial solo esta definido para numeros no negativos.")
		os.Exit(1)
	}	
	if numero > 20 {
		fmt.Printf("Advertencia: El factorial de %d es muy grande y probablemente cause un resultado incorrecto por desbordamiento (overflow).\n", numero)
	}
	resultado := factorialIterativo(numero)
	fmt.Printf("El factorial de %d es: %d\n", numero, resultado)
}