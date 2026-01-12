// fichero multiplo3.go
// Este programa dice si un número es múltiplo de 3 y a la vez par
package main
import ("bufio";"fmt";"os";"strconv";"strings")
func main() {
	var numero int64;var err error
	fmt.Print("Introduzca un número: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		texto := strings.TrimSpace(scanner.Text())
		numero, err = strconv.ParseInt(texto, 10, 64)
		if err != nil {fmt.Println("ERROR: Entrada no válida (no se introdujo un número)");os.Exit(1)}
		if (numero%3 == 0) && (numero%2 == 0) {
			fmt.Printf("\n El número %d es MÚLTIPLO DE 3 Y TAMBIÉN ES PAR (Múltiplo de 6) \n", numero)
		} else if numero%3 == 0 {fmt.Printf("El número %d ES MÚLTIPLO DE 3\n", numero)
          } else {fmt.Printf("El número %d NO ES MÚLTIPLO DE 3\n", numero)}
	}
}