// fichero pulgadasmilimetros.go
// Este programa cambia pulgadas por milímetros
package main
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func main() {
	const factorConversion = 25.4

	fmt.Print("Introduzca valor en pulgadas: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		texto := scanner.Text()
		texto = strings.Replace(texto, ",", ".", 1)
		pul, err := strconv.ParseFloat(texto, 64)
		if err != nil {
			fmt.Println("Error: El valor debe de ser un número.")
			os.Exit(1)
		}
		if pul < 0 {
			fmt.Println("Error: El valor debe ser un número positivo")
			os.Exit(1)
		}
		mil := factorConversion * pul
		fmt.Printf("El resultado en milímetros es: %v\n", mil)
	}
}