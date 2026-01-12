// fichero contarcara.go
// Ejemplo para contar caracteres y palabras leyendo una a una
package main
import ("bufio";"fmt";"io";"os";"unicode" )
func main() {
	fichero, err := os.Open("datos.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error al abrir fichero datos.txt: %v\n", err)
		os.Exit(1)
	}
	defer fichero.Close()
	// Usamos un Reader en lugar de Scanner para leer carácter a carácter
	reader := bufio.NewReader(fichero)
	contadorCaracteres := 0;contadorPalabras := 0;dentroDePalabra := false
	for {
		// ReadRune lee el siguiente carácter UTF-8
		r, _, err := reader.ReadRune()
		if err == io.EOF {break}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error leyendo el archivo: %v\n", err)
			os.Exit(1)
		}
		contadorCaracteres++
		if unicode.IsSpace(r) {
			dentroDePalabra = false
		} else {
			if !dentroDePalabra {
				contadorPalabras++
				dentroDePalabra = true
			}
		}
	}
	fmt.Printf("Total caracteres leídos: %d\n", contadorCaracteres)
	fmt.Printf("Total palabras: %d\n", contadorPalabras)
}