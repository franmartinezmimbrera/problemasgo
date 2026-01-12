// fichero areatriangulo1.go
// Este programa calcula el área de un triángulo rectángulo a partir de la base y la altura
package main
import ("bufio";"fmt";"os";"strconv";"strings")
func main() {
	var base, altura float64
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Introduce la base del triángulo rectángulo: ")
	if scanner.Scan() {
		texto := strings.Replace(scanner.Text(), ",", ".", 1)
		base, err = strconv.ParseFloat(texto, 64)
		if err != nil {fmt.Println("Error: El valor debe de ser un número.");os.Exit(1)}
		if base <= 0 {fmt.Println("Error: El valor debe ser un número > que 0");os.Exit(1)}
	}
	fmt.Print("Introduce la altura del triángulo rectángulo: ")
	if scanner.Scan() {
		texto := strings.Replace(scanner.Text(), ",", ".", 1)
		altura, err = strconv.ParseFloat(texto, 64)
		if err != nil {fmt.Println("Error: El valor debe de ser un número.");os.Exit(1)}
		if altura <= 0 {fmt.Println("Error: El valor debe ser un número > que 0");os.Exit(1)}
	}
	area := (base * altura) / 2.0
	fmt.Printf("El area del triángulo rectángulo es: %.2f\n", area)
}