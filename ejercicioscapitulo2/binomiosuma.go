// fichero binomiosuma.go
// Este programa calcula el binomio de suma de a+b al cuadrado
package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func main() {
	var a, b, resultado float64
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Introduce el valor de a: ")
	if scanner.Scan() {
		texto := strings.Replace(scanner.Text(), ",", ".", 1)
		a, err = strconv.ParseFloat(texto, 64)
		if err != nil {
			fmt.Println("Error: El valor debe ser un número.")
			os.Exit(1)
		}
	}
	fmt.Print("Introduce el valor de b: ")
	if scanner.Scan() {
		texto := strings.Replace(scanner.Text(), ",", ".", 1)
		b, err = strconv.ParseFloat(texto, 64)
		if err != nil {
			fmt.Println("Error: El valor debe ser un número.")
			os.Exit(1)
		}
	}
	resultado = (a * a) + (b * b) + (2.0 * a * b)
	fmt.Printf("El resultado del binomio de suma de (a+b)^2 es: %v\n", resultado)
}