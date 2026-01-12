// fichero cesar_cifrar.go
package main
import ("bufio";"fmt";"os";"strings";"unicode")
// cifrarCesar desplaza cada letra 'n' posiciones.
func cifrarCesar(texto string, n int) string {
	var sb strings.Builder
	shift := n % 26
	for _, r := range texto {
		if unicode.IsLetter(r) {	
			base := 'A'
			if unicode.IsLower(r) {base = 'a'}
			nuevo := base + (r-base+rune(shift))%26
			sb.WriteRune(nuevo)
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Introduce el mensaje a cifrar: ")
	mensaje, _ := reader.ReadString('\n')
	mensaje = strings.TrimSpace(mensaje)
	fmt.Print("Introduce la clave (número de desplazamiento): ")
	var n int
	fmt.Scanf("%d", &n)
	cifrado := cifrarCesar(mensaje, n)
	fmt.Printf("\nMensaje Cifrado: %s\n", cifrado)
}