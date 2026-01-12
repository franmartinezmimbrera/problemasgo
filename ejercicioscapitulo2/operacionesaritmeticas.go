// fichero operacionesaritmeticas.go
// Este programa realiza operaciones aritméticas dados 2 números
package main
import ("bufio";"fmt";"os";"strconv";"strings")
func main() {
	var valor1, valor2 float64; var err error
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Introduce el valor primero: ")
	if scanner.Scan() {
		texto1 := strings.Replace(scanner.Text(), ",", ".", 1); valor1, err = strconv.ParseFloat(texto1, 64)
		if err != nil { fmt.Println("ERROR: El valor introducido no es un número.");os.Exit(1)}
	}
	fmt.Print("Introduce el valor segundo: ")
	if scanner.Scan() {
		texto2 := strings.Replace(scanner.Text(), ",", ".", 1); valor2, err = strconv.ParseFloat(texto2, 64)
		if err != nil {fmt.Println("ERROR: El valor introducido no es un número.");os.Exit(1)}
	}
	fmt.Printf("El resultado de la suma (%v + %v) es: %v\n", valor1, valor2, valor1+valor2)
	fmt.Printf("El resultado de la resta (%v - %v) es: %v\n", valor1, valor2, valor1-valor2)
	fmt.Printf("El resultado de la multiplicación (%v * %v) es: %v\n", valor1, valor2, valor1*valor2)
	if valor2 == 0 {
		fmt.Printf("El resultado de la división (%v / %v) es: INDEFINIDO (División por Cero)\n", valor1, valor2)
	} else {
		fmt.Printf("El resultado de la división (%v / %v) es: %v\n", valor1, valor2, valor1/valor2)
	}
}