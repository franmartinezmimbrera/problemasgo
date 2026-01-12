// fichero alturatriangulo.go
// Este programa calcula la altura de un triángulo equilátero
package main
import ("bufio";"fmt";"math";"os";"strconv";"strings")

func main() {
	var l, h float64;var err error

	fmt.Print("Introduzca lado de un triángulo equilátero: ")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		texto := scanner.Text()
		texto = strings.Replace(texto, ",", ".", 1)
		l, err = strconv.ParseFloat(texto, 64)
		if err != nil {
			fmt.Println("Error: El valor debe ser un número.")
			os.Exit(1)
		}
		if l <= 0 {
			fmt.Println("Error: El lado de un triángulo equilátero debe ser un número mayor que 0.")
			os.Exit(1)
		}
		h = (math.Sqrt(3) * l) / 2.0
		fmt.Printf("La altura de un triángulo equilátero de lado %v es: %v\n", l, h)
	}
}