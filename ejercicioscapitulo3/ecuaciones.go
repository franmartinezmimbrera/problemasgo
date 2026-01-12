// fichero ecuaciones.go
// Resuelve ecuaciones de segundo grado por la fórmula general para números Reales
package main

import ("bufio";"fmt";"os";"strconv";"strings";"math")
func main() {
	var a, b, c, d, x1, x2 float64
	var err error
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Ingrese coeficiente a: ")	
		if !scanner.Scan() {os.Exit(0)}
		texto := strings.TrimSpace(scanner.Text())
		a, err = strconv.ParseFloat(texto, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: la entrada no es un número válido")
			os.Exit(1)
		}
		if a == 0 {
			fmt.Println("El coeficiente 'a' no puede ser cero.")
			continue 
		}
		break 
	}
	fmt.Print("Ingrese coeficiente b: ")
	if scanner.Scan() {
		b, err = strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: la entrada no es un número válido")
			os.Exit(1)
		}
	}
	fmt.Print("Ingrese coeficiente c: ")
	if scanner.Scan() {
		c, err = strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error: la entrada no es un número válido")
			os.Exit(1)
		}
	}
	d = math.Pow(b, 2) - 4*a*c
	if d > 0 {
		x1 = (-b + math.Sqrt(d)) / (2 * a)
		x2 = (-b - math.Sqrt(d)) / (2 * a)
		fmt.Printf("x1 = %f\n", x1)
		fmt.Printf("x2 = %f\n", x2)

	} else if d == 0 {
		x1 = -b / (2 * a)
		fmt.Printf("x1 = %f\n", x1)

	} else {
		fmt.Println("La ecuacion no tiene soluciones reales")
	}
}