// fichero ladotriangulo.go
// Este programa calcula el valor del lado a de un triángulo rectángulo
package main
import ("bufio";"fmt";"math";"os";"strconv";"strings")
func main() {
	var a, b, h float64;var err error
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Introduzca el valor del lado \"b\" del triángulo rectángulo: ")
	if scanner.Scan() {
		texto := strings.Replace(scanner.Text(), ",", ".", 1)
		b, err = strconv.ParseFloat(texto, 64)
		if err != nil {fmt.Println("Error: El valor debe ser un número.");os.Exit(1)}
		if b <= 0 {fmt.Println("Error: El lado 'b' debe ser un número mayor que 0.");os.Exit(1)}
	}
	fmt.Print("Introduzca el valor de la hipotenusa del triángulo rectángulo: ")
	if scanner.Scan() {
		texto := strings.Replace(scanner.Text(), ",", ".", 1)
		h, err = strconv.ParseFloat(texto, 64)		
		if err != nil {fmt.Println("Error: El valor debe ser un número.");os.Exit(1)}
		if h <= 0 {fmt.Println("Error: La hipotenusa debe ser un número mayor que 0.");os.Exit(1)}
	}
	if b >= h {
		fmt.Printf("Error: La hipotenusa (%v) debe ser mayor que el cateto 'b' (%v).\n", h, b)
		os.Exit(1)
	}
	a = math.Sqrt((h * h) - (b * b))
	fmt.Printf("El valor del lado a del triángulo rectángulo es: %v\n", a)
}