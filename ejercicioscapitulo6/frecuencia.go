//fichero frecuencia.go
package main
import ("fmt";"unicode")

func analizarFrecuencia(texto string) {
	conteo := make(map[rune]int)
	totalLetras := 0
	for _, r := range texto {
		if unicode.IsLetter(r) {
			r = unicode.ToUpper(r)
			conteo[r]++
			totalLetras++
		}
	}
	fmt.Printf("Análisis de '%s'...\n", texto)
	fmt.Printf("Total letras: %d\n\n", totalLetras)
	for letra := 'A'; letra <= 'Z'; letra++ {
		val := conteo[letra]
		porcentaje := (float64(val) / float64(totalLetras)) * 100
		if val > 0 {
			fmt.Printf("%c: %d (%.2f%%)\n", letra, val, porcentaje)
		}
	}
}
func main() {
	texto := "HOLA ESTO ES UNA PRUEBA DE FRECUENCIA EN GO LENGUAJE INTERESANTE"
	analizarFrecuencia(texto)
}