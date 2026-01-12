// fichero invierte.go
// Este ejercicio implementa la inversión de una cadena de texto dada
package main
import ("bufio";"fmt";"os")
// En Go, no pasamos arrays por referencia para modificarlos in-situ (normalmente),
// sino que devolvemos el nuevo valor transformado.
func invertirCadena(s string) string {
	// Convertimos el string a un slice de 'runes'.Un 'rune' es un carácter Unicode completo.
	// Si usáramos bytes simples, romperíamos caracteres con tildes o emojis (ej: 'ñ').
	runes := []rune(s)
	longitud := len(runes)
	inicio := 0
	fin := longitud - 1
	for inicio < fin {runes[inicio], runes[fin] = runes[fin], runes[inicio];inicio++;fin--}
	// Convertimos los runes de vuelta a string
	return string(runes)
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Introduce una cadena de texto: ")
	if scanner.Scan() {
		cadena := scanner.Text()
		if len(cadena) == 0 {fmt.Println("Cadena vacía.");os.Exit(0)}
		fmt.Printf("Cadena original:  \"%s\"\n", cadena)
		cadenaInvertida := invertirCadena(cadena)
		fmt.Printf("Cadena invertida: \"%s\"\n", cadenaInvertida)
	}
}