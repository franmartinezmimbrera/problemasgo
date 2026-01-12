// fichero estadisticas.go
// Este programa calcula estadísticas sobre alumnos
package main
import ("bufio";"fmt";"os";"strconv";"strings")
func main() {
	var suspensos, aprobados, notables, sobresalientes float64; var err error
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Introduce el número total de suspensos: ")
	if scanner.Scan() {
		texto := strings.Replace(scanner.Text(), ",", ".", 1)
		suspensos, err = strconv.ParseFloat(texto, 64)
		if err != nil {fmt.Println("Error: El valor debe de ser un número.");os.Exit(1)}
		if suspensos < 0 {fmt.Println("Error: El valor debe ser un número positivo");os.Exit(1)}
	}
	fmt.Print("Introduce el número total de aprobados: ")
	if scanner.Scan() {
		texto := strings.Replace(scanner.Text(), ",", ".", 1)
		aprobados, err = strconv.ParseFloat(texto, 64)
		if err != nil {fmt.Println("Error: El valor debe de ser un número.");os.Exit(1)}
		if aprobados < 0 {fmt.Println("Error: El valor debe ser un número positivo");os.Exit(1)}
	}
	fmt.Print("Introduce el número total de notables: ")
	if scanner.Scan() {
		texto := strings.Replace(scanner.Text(), ",", ".", 1)
		notables, err = strconv.ParseFloat(texto, 64)
		if err != nil {fmt.Println("Error: El valor debe de ser un número.");os.Exit(1)}
		if notables < 0 {fmt.Println("Error: El valor debe ser un número positivo");os.Exit(1)}
	}
	fmt.Print("Introduce el número total de sobresalientes: ")
	if scanner.Scan() {
		texto := strings.Replace(scanner.Text(), ",", ".", 1)
		sobresalientes, err = strconv.ParseFloat(texto, 64)
		if err != nil {fmt.Println("Error: El valor debe de ser un número.");os.Exit(1)}
		if sobresalientes < 0 {fmt.Println("Error: El valor debe ser un número positivo");os.Exit(1)}
	}
	totalalumnos := suspensos + aprobados + notables + sobresalientes
	if totalalumnos == 0 {
      fmt.Println("No se han introducido alumnos (Total 0), no se pueden calcular estadísticas.")
		os.Exit(1)
	}
	poraprot := ((aprobados + notables + sobresalientes) / totalalumnos) * 100
	porsus := (suspensos / totalalumnos) * 100
	pornot := (notables / totalalumnos) * 100
	porsobre := (sobresalientes / totalalumnos) * 100
	porapro := (aprobados / totalalumnos) * 100
	fmt.Printf("El porcentaje de alumnos que superan la asignatura es: %v\n", poraprot)
	fmt.Printf("El porcentaje de alumnos que suspenden la asignatura es: %v\n", porsus)
	fmt.Printf("El porcentaje de alumnos que sacan notable en la asignatura es: %v\n", pornot)
	fmt.Printf("El porcentaje de alumnos que sobresalen en la asignatura es: %v\n", porsobre)
	fmt.Printf("El porcentaje de alumnos que han sacado un aprobado en la asignatura es: %v\n", porapro)
}