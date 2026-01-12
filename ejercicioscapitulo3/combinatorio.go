// fichero combinatorio.go
// Calcula el número combinatorio C(n, r) o n sobre r, usando factoriales.
package main
import ("bufio";"fmt";"os";"strconv";"strings")
// Función factorial iterativa (igual que en el ejercicio anterior)
func factorial(n int64) int64 {
	if n == 0 || n == 1 {return 1}
	resultado := int64(1)
	for i := int64(2); i <= n; i++ {resultado *= i}
	return resultado
}
// Función para calcular un número combinatorio nCr
func nCr(n, r int64) int64 {
	// 1. Validaciones básicas de rango
	if r < 0 || r > n {return 0}
	// 2. Cálculo de los tres factoriales necesarios
	factN := factorial(n)
	factR := factorial(r)
	factNMinusR := factorial(n - r)
	// Si algún factorial dio negativo, es que "dio la vuelta" al int64 (superó el límite).
	if factN < 0 || factR < 0 || factNMinusR < 0 {
		fmt.Fprintln(os.Stderr, "Error: Un factorial intermedio es demasiado grande para 'int64'.")
		return -1 // Código de error
	}
	return factN / (factR * factNMinusR)
}
func main() {
	var n, m int64;var err error
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Introduce un numero entero no negativo para N (max 20 para precision): ")
	if scanner.Scan() {
		texto := strings.TrimSpace(scanner.Text())
		n, err = strconv.ParseInt(texto, 10, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: Entrada no válida para N.")
			os.Exit(1)
		}
	} else {
		return
	}
	if n < 0 {
		fmt.Fprintln(os.Stderr, "Error: N debe ser no negativo.");os.Exit(1)}
	if n > 20 {
		fmt.Printf("Advertencia: El factorial de N=%d causará desbordamiento.\n", n)
	}
	fmt.Print("Introduce un numero entero no negativo para M (max 20 para precision): ")
	if scanner.Scan() {
		texto := strings.TrimSpace(scanner.Text())
		m, err = strconv.ParseInt(texto, 10, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: Entrada no válida para M.")
			os.Exit(1)
		}
	} else {return}
	if m < 0 {fmt.Fprintln(os.Stderr, "Error: M debe ser no negativo.");os.Exit(1)}
	if m > 20 {fmt.Printf("Advertencia: El factorial de M=%d causará desbordamiento.\n", m)}
	combinaciones := nCr(n, m)
	if combinaciones > 0 {
		fmt.Printf("El número combinatorio de %d sobre %d es: %d\n", n, m, combinaciones)
	} else if combinaciones == 0 {
		fmt.Printf("El número combinatorio de %d sobre %d es 0 (M fuera de rango 0 <= M <= N).\n", n, m)
	} else if combinaciones == -1 {
		fmt.Printf("No se pudo calcular el número combinatorio de %d sobre %d debido a desbordamiento.\n", n, m)
	}
}