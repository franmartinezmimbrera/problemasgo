// fichero tipotriangulo.go
// Este programa calcula el tipo de triángulo en función de los lados
package main
import ("bufio";"fmt";"os";"strconv";"strings")
func main() {
	var l1, l2, l3 float64; var err error
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Introduce lo que mide el primer lado del triángulo: ")
	if scanner.Scan() {
		l1, err = strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
		if err != nil {fmt.Println("Error: la entrada no es un número válido");os.Exit(1)}
	}
	fmt.Print("Introduce lo que mide el segundo lado del triángulo: ")
	if scanner.Scan() {
		l2, err = strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
		if err != nil {fmt.Println("Error: la entrada no es un número válido");os.Exit(1)}
	}
	fmt.Print("Introduce lo que mide el tercer lado del triángulo: ")
	if scanner.Scan() {
		l3, err = strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
		if err != nil {fmt.Println("Error: la entrada no es un número válido");os.Exit(1)}
	}
	if l1+l2 <= l3 || l1+l3 <= l2 || l2+l3 <= l1 {
        fmt.Println("\nError Geométrico: Los lados NO forman un triángulo válido.");os.Exit(1)
	}
	if l1 == l2 && l2 == l3 {fmt.Println("El Triangulo es Equilatero")
	} else if l1 == l2 || l1 == l3 || l2 == l3 {
		fmt.Println("El Triangulo es Isoceles")
	} else {fmt.Println("El Triangulo es Escaleno")}
}