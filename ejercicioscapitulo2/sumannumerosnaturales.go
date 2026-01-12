// fichero sumannumerosnaturales.go
// Este programa calcula la suma de los "n" primeros números naturales
package main
import (
	"bufio";"fmt";"os";"strconv"
)
func main() {
	var n int
	var suma int64 
	fmt.Print("Introduzca número de números naturales a sumar: ")

	scanner := bufio.NewScanner(os.Stdin)	
	if scanner.Scan() {
		texto := scanner.Text()
		val, err := strconv.Atoi(texto)
		if err != nil {
			fmt.Println("Error: El valor debe de ser un número.")
			os.Exit(1)
		}
		n = val
		if n <= 0 {
			fmt.Println("Error: El valor debe ser un número natural.")
			os.Exit(1)
		}
		suma = (int64(n) * (int64(n) + 1)) / 2
		fmt.Printf("La suma de los %d primeros números naturales es: %d\n", n, suma)
	}
}