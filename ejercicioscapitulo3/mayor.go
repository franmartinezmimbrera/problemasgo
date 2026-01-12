// fichero mayor.go
// Este programa calcula el número mayor de 10 introducidos por teclado
package main
import ("bufio";"fmt";"os";"strconv";"strings")
const NumElementos = 10
func mayor(numeros []int64) int64 {
	if len(numeros) == 0 {return 0}
	max := numeros[0]
	for _, num := range numeros {
		if num > max {max = num}
	}
	return max
}
func main() {
	var numerosin []int64; var err error
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < NumElementos; i++ {
		for {
			fmt.Printf("Introduzca número %d de %d: ", i+1, NumElementos)
			if !scanner.Scan() {os.Exit(0)}
			texto := strings.TrimSpace(scanner.Text())
			numeroLeido, err := strconv.ParseInt(texto, 10, 64)
			if err != nil {fmt.Println("ERROR: La entrada no es un número entero válido. Inténtelo de nuevo");continue}			
			numerosin = append(numerosin, numeroLeido)			
			break 
		}
	}
	fmt.Printf("\nEl número mayor de todos los introducidos es: %d\n", mayor(numerosin))
}