// fichero areatriangulo2.go 
// Este programa calcula el área de un triángulo con la fórmula de Herón
package main
import ("bufio";"fmt";"os";"strconv";"strings";"math")
func AreaTrianguloHeron(l1, l2, l3 float64) float64 {
	if l1 <= 0 || l2 <= 0 || l3 <= 0 || l1+l2 <= l3 || l1+l3 <= l2 || l2+l3 <= l1 {
		return 0.0
	}
	sp := (l1 + l2 + l3) / 2.0
	// Fórmula de Herón: Raíz(s * (s-a) * (s-b) * (s-c))
	area := math.Sqrt(sp * (sp - l1) * (sp - l2) * (sp - l3))
	return area
}
func main() {
	var l1, l2, l3 float64;var err error
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Introduce lo que mide el primer lado del triángulo: ")
	if scanner.Scan() {
		l1, err = strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
		if err != nil {fmt.Println("ERROR: Entrada no válida (no es un número).");os.Exit(1)}
	}
	fmt.Print("Introduce lo que mide el segundo lado del triángulo: ")
	if scanner.Scan() {
		l2, err = strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
		if err != nil {fmt.Println("ERROR: Entrada no válida (no es un número).");os.Exit(1)}
	}
	fmt.Print("Introduce lo que mide el tercer lado del triángulo: ")
	if scanner.Scan() {
		l3, err = strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
		if err != nil {fmt.Println("ERROR: Entrada no válida (no es un número).");os.Exit(1)}
	}
	areaFinal := AreaTrianguloHeron(l1, l2, l3)
	if areaFinal > 0.0 {
		fmt.Printf("El área del triángulo es: %.2f\n", areaFinal)
	} else {
		fmt.Println("Los lados introducidos NO forman un triángulo válido (la suma de 2 lados debe ser > que el tercero, y todos deben ser >=0)")
	}
}