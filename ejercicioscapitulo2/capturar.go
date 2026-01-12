// fichero capturar.go
// Captura información desde teclado y la muestra
package main
import (
	"bufio";"fmt";"os";"strconv";"strings" 
)
func main() {
	var edad int ; var peso float64 ;var err error
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Dinos tu Edad, peso y color favorito:")
	fmt.Print("\n       Edad: ")
	if scanner.Scan() {
		textoEdad := scanner.Text()
		edad, err = strconv.Atoi(textoEdad)
		if err != nil {fmt.Println("Error: La edad debe ser un número entero."); os.Exit(1) }
	}
	fmt.Print("\n       Peso: ")
	if scanner.Scan() {
		textoPeso := scanner.Text()
		textoPeso = strings.Replace(textoPeso, ",", ".", 1)
		peso, err = strconv.ParseFloat(textoPeso, 64)
		if err != nil {
			fmt.Println("Error: El peso debe ser un número.")
			os.Exit(1)
		}
	}
	fmt.Print("\n Color favorito: ")
	scanner.Scan()
	color := scanner.Text()
	fmt.Printf("El %s!!! \n", color)
	fmt.Printf("¿Cómo puede gustarte el %s con %d años y pesando %.2fKg.?\n", color, edad, peso)
}