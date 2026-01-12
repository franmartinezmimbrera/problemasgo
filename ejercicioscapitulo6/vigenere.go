//vigenere.go
package main
import ("bufio";"fmt";"os";"strings";"unicode")
func vigenere(texto, clave string, cifrar bool) string {
	var sb strings.Builder
	claveRunes := []rune(strings.ToUpper(clave)) 
	kLen := len(claveRunes)
	kIndex := 0 // Índice para recorrer la clave 
	for _, r := range texto {
		if unicode.IsLetter(r) {
			base := 'A'
			if unicode.IsLower(r) {base = 'a'}			
			valTexto := int(r - base)
			valClave := int(claveRunes[kIndex%kLen] - 'A')
			var nuevoVal int
			if cifrar {
				nuevoVal = (valTexto + valClave) % 26
			} else {
				nuevoVal = (valTexto - valClave + 26) % 26
			}
			sb.WriteRune(base + rune(nuevoVal))
			kIndex++ 
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Mensaje: ")
	msg, _ := reader.ReadString('\n')
	msg = strings.TrimSpace(msg)
	fmt.Print("Clave: ")
	key, _ := reader.ReadString('\n')
	key = strings.TrimSpace(key)
	cifrado := vigenere(msg, key, true)
	fmt.Printf("Cifrado: %s\n", cifrado)
	descifrado := vigenere(cifrado, key, false)
	fmt.Printf("Descifrado: %s\n", descifrado)
}