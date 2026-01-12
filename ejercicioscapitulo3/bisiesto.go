// fichero bisiesto.go
// Este programa dice si un año es bisiesto o no
package main
import ("bufio";"fmt";"os";"strconv";"strings")
func main() {
	var anio int64;var err error
	fmt.Print("Introduzca un año: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		texto := strings.TrimSpace(scanner.Text())
		anio, err = strconv.ParseInt(texto, 10, 64)
		if err != nil {fmt.Println("ERROR: La entrada no es un número entero");os.Exit(1)}
		if (anio%4 == 0 && anio%100 != 0) || (anio%400 == 0) {fmt.Println("ES BISIESTO")
		} else {fmt.Println("NO ES BISIESTO")}
	}
}