// fichero leelongi.go
// Este ejercicio lee una cadena desde teclado, dice la longitud y la concatena con otra.
package main
import ("bufio";"fmt";"os";"unicode/utf8")

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Introduzca cadena (o solo Enter para terminar): ")
		if !scanner.Scan() {
			break 
		}
		cadena := scanner.Text()
		if len(cadena) == 0 {
			break
		}
		lenBytes := len(cadena)
		lenCaracteres := utf8.RuneCountInString(cadena)
		fmt.Printf("La cadena contiene %d bytes y %d caracteres reales.\n", lenBytes, lenCaracteres)
		cadena += ".txt"	
		fmt.Printf("Concatenación: %s \n", cadena)
	}
}