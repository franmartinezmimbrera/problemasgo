// fichero parimpar.go
// Este programa dice si un número es par o impar
package main
import ("bufio";"fmt";"os";"strconv";"strings")
func main() {
	var numero int64;var err error
	fmt.Print("Introduzca un número entero: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		texto := strings.TrimSpace(scanner.Text())
		numero, err = strconv.ParseInt(texto, 10, 64)
		if err != nil {
			fmt.Println("Entrada no válida (no se introdujo un número entero)")
			os.Exit(1)
		}
		if numero%2 == 0 {
			fmt.Println("ES PAR")
		} else {
			fmt.Println("ES IMPAR")
		}
	}
}
