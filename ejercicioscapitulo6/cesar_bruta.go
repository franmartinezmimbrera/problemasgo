// fichero cesar_bruta.go
package main
import ("bufio";"fmt";"os";"strings";"unicode")
func descifrar(texto string, n int) string {
	var sb strings.Builder
	shift := n % 26
	for _, r := range texto {
		if unicode.IsLetter(r) {
			base := 'A'
			if unicode.IsLower(r) {base = 'a'}
			tmp := (r - base - rune(shift)) % 26
			if tmp < 0 { tmp += 26 }
			sb.WriteRune(base + tmp)
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Introduce el mensaje cifrado para atacar: ")
	mensaje, _ := reader.ReadString('\n')
	mensaje = strings.TrimSpace(mensaje)
	fmt.Println("\n Iniciando Ataque de Fuerza Bruta ")
	for k := 1; k < 26; k++ {
		candidato := descifrar(mensaje, k)
		fmt.Printf("Clave %2d: %s\n", k, candidato)
	}
	fmt.Println("Fin del Ataque")
	fmt.Println("Busca visualmente la frase que tenga sentido.")
}