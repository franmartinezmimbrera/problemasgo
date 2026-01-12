// fichero volumencilindro.go
// Calculamos el volumen de un cilindro
package main
import ("bufio";"fmt";"math";"os";"strconv";"strings")
func main() {
	var D, H, R, V float64; var err error
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Introduzca el diámetro, en metros: ")
	if scanner.Scan() {
		texto := strings.Replace(scanner.Text(), ",", ".", 1)
		D, err = strconv.ParseFloat(texto, 64)
		if err != nil {
			fmt.Println("Error: El valor debe ser un número.")
			os.Exit(1)
		}
	}
	fmt.Print("Introduzca la altura, en metros: ")
	if scanner.Scan() {
		texto := strings.Replace(scanner.Text(), ",", ".", 1)
		H, err = strconv.ParseFloat(texto, 64)
		if err != nil {
			fmt.Println("Error: El valor debe ser un número.")
			os.Exit(1)
		}
	}
	if H <= 0 || D <= 0 {
		fmt.Println("Error: El diámetro y la altura debe ser un número > que 0")
		os.Exit(1)
	}
    R = D / 2.0;V = math.Pi * R * R * H
	fmt.Printf("El volumen del cilindro es de %v m^3\n", V)
}