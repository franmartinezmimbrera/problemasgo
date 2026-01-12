// fichero newton.go
// Calcula mediante el binomio de newton los datos dados
package main
import ("bufio";"fmt";"os";"strconv";"strings")

// Función factorial iterativa
func factorial(n int64) int64 {
	if n == 0 || n == 1 {return 1}
	resultado := int64(1)
	for i := int64(2); i <= n; i++ {resultado *= i}
	return resultado
}

// Función para calcular un número combinatorio (nCr)
func nCr(n, r int64) int64 {
	if r < 0 || r > n {return 0}
	factN := factorial(n)
	factR := factorial(r)
	factNMinusR := factorial(n - r)
	if factN < 0 || factR < 0 || factNMinusR < 0 {
		fmt.Fprintln(os.Stderr, "Error: Un factorial intermedio es demasiado grande para 'int64'.")
		return -1
	}
	return factN / (factR * factNMinusR)
}

// calcularPotencia implementa "Exponenciación Binaria"
// Es un algoritmo mucho más eficiente que multiplicar en bucle uno a uno.
func calcularPotencia(base, exponente int64) int64 {
	resultado := int64(1)
	for exponente > 0 {
		if exponente%2 == 1 {
			resultado *= base
		}
		base *= base
		exponente /= 2
	}
	return resultado
}

// binomioDeNewton calcula la expansión (a + b)^n
func binomioDeNewton(a, b, n int64) {
	if n < 0 || n > 20 {
		fmt.Fprintln(os.Stderr, "Error: El exponente N debe ser no negativo y <= 20.")
		return
	}
	fmt.Printf("\nExpansión de (%d + %d)^%d:\n", a, b, n)
	var suma int64 = 0
	for k := int64(0); k <= n; k++ {
		coeficiente := nCr(n, k)
		if coeficiente == -1 {
			fmt.Fprintln(os.Stderr, "Cálculo abortado debido a un error de factorial.")
			return
		}
		potenciaA := calcularPotencia(a, n-k)
		potenciaB := calcularPotencia(b, k)
		valorTermino := coeficiente * potenciaA * potenciaB
		suma += valorTermino
		fmt.Printf("  + (%d * %d^%d * %d^%d) = %d\n", 
			coeficiente, a, n-k, b, k, valorTermino)
	}
	fmt.Printf("Valor Total del Binomio de Newton de (a+b)^n: %d\n", suma)
}

func main() {
	var a, b, n int64;var err error
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Introduce un numero entero no negativo para a: ")
	if scanner.Scan() {
		a, err = strconv.ParseInt(strings.TrimSpace(scanner.Text()), 10, 64)
		if err != nil || a < 0 {
			fmt.Fprintln(os.Stderr, "Error: Entrada no válida para a (debe ser no negativa).")
			os.Exit(1)
		}
	}
	fmt.Print("Introduce un numero entero no negativo para b: ")
	if scanner.Scan() {
		b, err = strconv.ParseInt(strings.TrimSpace(scanner.Text()), 10, 64)
		if err != nil || b < 0 {
			fmt.Fprintln(os.Stderr, "Error: Entrada no válida para b (debe ser no negativa).")
			os.Exit(1)
		}
	}
	fmt.Print("Introduce un numero entero no negativo para n: ")
	if scanner.Scan() {
		n, err = strconv.ParseInt(strings.TrimSpace(scanner.Text()), 10, 64)
		if err != nil || n < 0 {
			fmt.Fprintln(os.Stderr, "Error: Entrada no válida para n (debe ser no negativo).")
			os.Exit(1)
		}
	}
	// Nota: (a > 1 && n > 15) es una heurística simple para avisar que el número será gigante
	if n > 20 || (a > 1 && n > 15) || (b > 1 && n > 15) {
		fmt.Println("\nAdvertencia: La expansión de (a+b)^n es muy grande. El resultado total puede desbordar 'int64'.")
	}
	binomioDeNewton(a, b, n)
}