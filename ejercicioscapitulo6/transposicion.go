// transposicion.cpp 
package main
import ("fmt";"math";"strings")

func cifrarTransposicion(mensaje string, ancho int) string {
	// Limpiamos espacios para el ejemplo clásico de Escítala
	mensaje = strings.ReplaceAll(mensaje, " ", "")
	runes := []rune(mensaje)
	n := len(runes)
	// Calculamos la altura necesaria (redondeo hacia arriba)
	filas := int(math.Ceil(float64(n) / float64(ancho)))
	var sb strings.Builder
	for col := 0; col < ancho; col++ {
		for fila := 0; fila < filas; fila++ {
			index := fila*ancho + col
			if index < n {
				sb.WriteRune(runes[index])
			}
		}
	}
	return sb.String()
}
func main() {
	msg := "ESTO ES UN EJEMPLO DE TRANSPOSICION"
	ancho := 4 // Clave
	fmt.Printf("Original: %s\n", msg)
	fmt.Printf("Ancho (Clave): %d\n", ancho)	
	cifrado := cifrarTransposicion(msg, ancho)
	fmt.Printf("Cifrado:  %s\n", cifrado)
}