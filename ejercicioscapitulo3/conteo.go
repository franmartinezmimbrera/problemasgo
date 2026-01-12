// fichero conteo.go
// Este ejercicio convierte cadenas a mayúsculas y cuenta sus vocales.
package main
import ("bufio";"fmt";"os";"strings")
const numFrases = 4
func convertirAMayusculas(s string) string {
	return strings.ToUpper(s)
}
func contarVocales(s string) int {
	contador := 0
	// 'range' descompone el string en 'runes' (caracteres Unicode seguros).
	// Si usáramos un for clásico (i=0; i<len), estaríamos leyendo bytes, no letras.
	for _, letra := range s {
		// En Go, el switch no necesita 'break'.
		switch letra {
		case 'A', 'E', 'I', 'O', 'U':
			contador++
		}
	}
	return contador
}

func main() {
	listaFrases := make([]string, numFrases)
	scanner := bufio.NewScanner(os.Stdin)
	totalVocales := 0
	fmt.Printf("Introduce %d frases/líneas de texto:\n", numFrases)
	for i := 0; i < numFrases; i++ {
		fmt.Printf("Frase %d: ", i+1)
		if scanner.Scan() {
			listaFrases[i] = scanner.Text()
		} else {
			fmt.Fprintln(os.Stderr, "Error en la lectura o fin de archivo.")
			os.Exit(1)
		}
	}
	fmt.Println("\n Procesamiento y Conteo ")
	for i := 0; i < numFrases; i++ {
		listaFrases[i] = convertirAMayusculas(listaFrases[i])
		vocalesFrase := contarVocales(listaFrases[i])
		totalVocales += vocalesFrase
		fmt.Printf("Frase %d (MAYÚS): '%s' -> Vocales contadas: %d\n", 
			i+1, listaFrases[i], vocalesFrase)
	}
	fmt.Println("\n Resumen")
	fmt.Printf("El número total de vocales en todas las frases es: %d\n", totalVocales)
}