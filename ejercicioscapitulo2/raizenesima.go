// fichero raizenesima.go
// Este programa calcula la raíz n-ésima de un número
package main
import ("bufio";"fmt":"math";"os";"strconv";"strings")
func main() {
	var numero, resultado float64;var exponente int;var err error
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Introduce el número a calcular la raíz: ")
	if scanner.Scan() {
		texto := strings.Replace(scanner.Text(), ",", ".", 1)
		numero, err = strconv.ParseFloat(texto, 64)
		if err != nil {fmt.Println("Error: El valor debe de ser un número.");os.Exit(1)}
	}
	fmt.Print("Introduce exponente de raíz (un entero, ej: 2, 3..): ")
	if scanner.Scan() {
		texto := scanner.Text()
		exponente, err = strconv.Atoi(texto)
		if err != nil {fmt.Println("Error: El valor debe de ser un número entero.");os.Exit(1)}
	}
	if exponente == 0 {fmt.Println("Error: El exponente de la raíz no puede ser 0");os.Exit(1)}
	if numero < 0 && exponente%2 == 0 {
		fmt.Println("Imposible cálculo de raíz par de un número < 0 en R")
		os.Exit(1)
	}
	basePositiva := math.Abs(numero)
	potencia := 1.0 / float64(exponente)	
	resultado = math.Pow(basePositiva, potencia)
	if numero < 0 && exponente%2 != 0 { resultado = -resultado}
	fmt.Printf("La raíz %d ésima de %v es: %v\n", exponente, numero, resultado)
}