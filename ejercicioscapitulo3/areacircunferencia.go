// fichero areacircunferencia.go
// Este programa calcula el Área de una circunferencia
package main
import ("bufio";"fmt";"os";"strconv";"strings";"math")
func AreaCircunferencia(r float64) float64 {
	return math.Pi * math.Pow(r, 2)
}
func main() {
	var radio float64;var err error
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Introduce el radio de la circunferencia >0: ")
		if !scanner.Scan() {break }
		texto := strings.TrimSpace(scanner.Text())
		radio, err = strconv.ParseFloat(texto, 64)
		if err != nil {fmt.Println("ERROR: La entrada no es un número válido.");continue }
		if radio <= 0 {fmt.Println("ERROR: El radio debe ser mayor que cero");continue }
		break
	}
	area := AreaCircunferencia(radio)
	fmt.Printf("El área de la circunferencia es: %f\n", area)
}