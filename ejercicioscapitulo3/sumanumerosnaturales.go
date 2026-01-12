// fichero sumanaturales.go
// Este programa calcula la suma de los "n" primeros números naturales con for
package main
import ("bufio";"fmt";"os";"strconv";"strings")
func main() {
	var numero int64;var suma int64; var err error
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Introduzca un número n >0 de términos a sumar: ")
		if !scanner.Scan() {break}

		texto := strings.TrimSpace(scanner.Text())
		numero, err = strconv.ParseInt(texto, 10, 64)
		if err != nil {
			fmt.Println("ERROR: La entrada no es un número entero válido.")
			continue
		}
		if numero <= 0 {
			fmt.Println("ERROR: El número de términos debe ser >0.")
			continue
		}
		break
	}
	for i := int64(1); i <= numero; i++ {
		suma = suma + i
	}
	fmt.Printf("La suma de los %d primeros números naturales es: %d\n", numero, suma)
}