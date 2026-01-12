// fichero sumanumeros.go
// Este programa calcula la suma de los números introducidos hasta leer -50
package main
import ("bufio";"fmt";"os";"strconv";"strings")
func main() {
	var numero, suma float64;var err error
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Introduzca número a sumar (o -50 para terminar): ")
		if !scanner.Scan() {break }
		texto := strings.TrimSpace(scanner.Text())
		numero, err = strconv.ParseFloat(texto, 64)
		if err != nil {
			fmt.Println("ERROR: La entrada no es un número válido. Inténtelo de nuevo.")
			continue
		}
		if numero == -50.0 {
			break 
		}
		suma += numero
	}
	fmt.Printf("La suma total de los números introducidos es: %f\n", suma)
}