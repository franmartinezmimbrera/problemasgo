// fichero areatriangulo2.go
// Este programa calcula el Área de un triángulo equilátero a partir de uno de sus lados
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
	var lado float64
	var err error
	fmt.Print("Introduce un lado del triángulo equilátero: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		texto := scanner.Text()
		texto = strings.Replace(texto, ",", ".", 1)
		lado, err = strconv.ParseFloat(texto, 64)
		if err != nil {
			fmt.Println("Error: El valor debe ser un número.")
			os.Exit(1)
		}
		if lado <= 0 {
			fmt.Println("Error: El valor debe ser un número mayor que 0.")
			os.Exit(1)
		}
		area := (math.Sqrt(3.0) / 4.0) * lado * lado
		fmt.Printf("El Área del triángulo equilátero de lado %v es: %v\n", lado, area)
	}
}