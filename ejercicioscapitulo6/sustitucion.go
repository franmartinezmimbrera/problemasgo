//Fichero sustitucion.go
package main
import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "unicode"
)

// Alfabeto desordenado de ejemplo (Clave)
const alfabetoNormal = "abcdefghijklmnopqrstuvwxyz"
const alfabetoClave  = "qwertyuiopasdfghjklzxcvbnm"

func sustitucion(texto string, mapeo map[rune]rune) string {

    var sb strings.Builder
	for _, r := range texto {
		// Convertimos a minúscula para buscar en el mapa
		lower := unicode.ToLower(r)	
		if val, ok := mapeo[lower]; ok {
			// Si es mayúscula, convertimos el resultado a mayúscula
			if unicode.IsUpper(r) {
				sb.WriteRune(unicode.ToUpper(val))
			} else {
				sb.WriteRune(val)
			}
		} else {
			sb.WriteRune(r)
		}
	}
	return sb.String()
}

func main() {

	mapero := make(map[rune]rune)
	runesNormal := []rune(alfabetoNormal)
	runesClave := []rune(alfabetoClave)

    for i := 0; i < 26; i++ {
		mapero[runesNormal[i]] = runesClave[i]
	}

    reader := bufio.NewReader(os.Stdin)
	fmt.Println("Clave de sustitución: ", alfabetoClave)
	fmt.Print("Mensaje a cifrar: ")
	msg, _ := reader.ReadString('\n')	
	resultado := sustitucion(strings.TrimSpace(msg), mapero)
	fmt.Printf("Resultado: %s\n", resultado)
}