// fichero cesar_descifrar.go
package main
import ("bufio";"fmt";"os";"strings";"unicode")
func descifrarCesar(texto string, n int) string {
	var sb strings.Builder
	shift := n % 26
	for _, r := range texto {
		if unicode.IsLetter(r) {
			base := 'A'
			if unicode.IsLower(r) {base = 'a'}
			tmp := (r - base - rune(shift)) % 26
			if tmp < 0 {tmp += 26}
			nuevo := base + tmp
			sb.WriteRune(nuevo)
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Introduce el mensaje cifrado: ")
	mensaje, _ := reader.ReadString('\n')
	mensaje = strings.TrimSpace(mensaje)
	fmt.Print("Introduce la clave original: ")
	var n int
	fmt.Scanf("%d", &n)
	claro := descifrarCesar(mensaje, n)
	fmt.Printf("\nMensaje Descifrado: %s\n", claro)
}