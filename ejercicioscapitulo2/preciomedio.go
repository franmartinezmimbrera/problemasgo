// fichero preciomedio.go
// Este programa calcula la media de tres precios
package main
import ("bufio";"fmt";"os";"strconv";"strings")
func main() {
	var precio1, precio2, precio3 float64;var err error
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Introduzca el precio en establecimiento 1, en euros: ")
	if scanner.Scan() {
		texto := strings.Replace(scanner.Text(), ",", ".", 1); precio1, err = strconv.ParseFloat(texto, 64)		
		if err != nil {fmt.Println("Error: El valor debe de ser un número.");os.Exit(1)}
		if precio1 < 0 {fmt.Println("Ningún precio puede ser negativo");os.Exit(1)}}
	fmt.Print("Introduzca el precio en establecimiento 2, en euros: ")
	if scanner.Scan() {
		texto := strings.Replace(scanner.Text(), ",", ".", 1);precio2, err = strconv.ParseFloat(texto, 64)
		if err != nil {fmt.Println("Error: El valor debe de ser un número.");os.Exit(1)}
		if precio2 < 0 {fmt.Println("Ningún precio puede ser negativo");os.Exit(1)}}
	fmt.Print("Introduzca el precio en establecimiento 3, en euros: ")
	if scanner.Scan() {
		texto := strings.Replace(scanner.Text(), ",", ".", 1);precio3, err = strconv.ParseFloat(texto, 64)
		if err != nil {fmt.Println("Error: El valor debe de ser un número.");os.Exit(1)}
		if precio3 < 0 {fmt.Println("Ningún precio puede ser negativo");os.Exit(1)}}
	media := (precio1 + precio2 + precio3) / 3.0
	fmt.Printf("El precio medio del producto es de %v euros\n", media)
}
