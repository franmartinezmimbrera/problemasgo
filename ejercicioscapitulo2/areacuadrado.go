// fichero areacuadrado.go
// Este programa calcula el área de un cuadrado a partir de uno de sus lados
package main
import (
	"bufio";"fmt";"os";"strconv";"strings"
)
func main() {
	fmt.Print("Introduce valor lado de un cuadrado: ")
	scanner := bufio.NewScanner(os.Stdin)	
	if scanner.Scan() {
		textoLado := scanner.Text()
		textoLado = strings.Replace(textoLado, ",", ".", 1)
		l1, err := strconv.ParseFloat(textoLado, 64)
		if err != nil {
			fmt.Println("Error: El lado debe de ser un número.")
			os.Exit(1)
		}
		if l1 <= 0 {
			fmt.Println("Error: El lado debe ser un número positivo >0")
			os.Exit(1)
		}
		area := l1 * l1
		fmt.Printf("El Área del cuadrado es: %v \n", area)
	}
}