// fichero perimetrocir.go
// Este programa calcula el perímetro de una circunferencia
package main

import ("bufio";"fmt";"math";"os";"strconv";"strings")
func main() {
	fmt.Print("Introduzca el radio: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		texto := scanner.Text()
		texto = strings.Replace(texto, ",", ".", 1)
		radio, err := strconv.ParseFloat(texto, 64)
		if err != nil {
			fmt.Println("Error: El valor debe de ser un número.")
			os.Exit(1)
		}
		if radio <= 0 {
			fmt.Println("Error: El radio debe ser un número >=0")
			os.Exit(1)
		}
		perimetro := 2 * math.Pi * radio		
		fmt.Printf("El perímetro es: %v\n", perimetro)
	}
}