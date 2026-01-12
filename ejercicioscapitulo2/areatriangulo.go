// fichero areatriangulo.go
// Calculamos el área de un triángulo mediante la fórmula de Herón
package main
import (
    "bufio"
    "fmt"
    "math"
    "os"
    "strconv"
    "strings"
)
func main() {
	var l1, l2, l3 float64
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Introduce lo que mide el primer lado del triángulo: ")
	if scanner.Scan() {
		texto := strings.Replace(scanner.Text(), ",", ".", 1)
		l1, err = strconv.ParseFloat(texto, 64)
		if err != nil {
			fmt.Println("Error: El valor debe de ser un número.")
			os.Exit(1)
		}
		if l1 < 0 {
			fmt.Println("Error: El valor debe ser un número positivo")
			os.Exit(1)
		}
	}
	fmt.Print("Introduce lo que mide el segundo lado del triángulo: ")
	if scanner.Scan() {
		texto := strings.Replace(scanner.Text(), ",", ".", 1)
		l2, err = strconv.ParseFloat(texto, 64)
		if err != nil {
			fmt.Println("Error: El valor debe de ser un número.")
			os.Exit(1)
		}
		if l2 < 0 {
			fmt.Println("Error: El valor debe ser un número positivo")
			os.Exit(1)
		}
	}
	fmt.Print("Introduce lo que mide el tercer lado del triángulo: ")
	if scanner.Scan() {
		texto := strings.Replace(scanner.Text(), ",", ".", 1)
		l3, err = strconv.ParseFloat(texto, 64)
		if err != nil {
			fmt.Println("Error: El valor debe de ser un número.")
			os.Exit(1)
		}
		if l3 < 0 {
			fmt.Println("Error: El valor debe ser un número positivo")
			os.Exit(1)
		}
	}
	if l1+l2 <= l3 || l1+l3 <= l2 || l2+l3 <= l1 {
		fmt.Printf("\nError: Las longitudes (%v, %v, %v) no forman un triángulo válido (desigualdad triangular).\n", l1, l2, l3)
		os.Exit(1)
	}
	sp := (l1 + l2 + l3) / 2.0
	area := math.Sqrt(sp * (sp - l1) * (sp - l2) * (sp - l3))
	fmt.Printf("El area del triángulo es: %.4f\n", area)
}