//fichero procesar_csv.go
package main
import ("encoding/csv";"fmt";"os";"strconv")

func main() {
	// Creamos un CSV de ejemplo en memoria
	archivoNombre := "datos.csv"
	f, _ := os.Create(archivoNombre)
	f.WriteString("Nombre,Edad,Nota\nAna,22,8.5\nLuis,24,9.0\nPepe,21,4.5")
	f.Close()
	// Abrimos para leer
	file, _ := os.Open(archivoNombre)
	defer file.Close()
	lector := csv.NewReader(file)
	registros, _ := lector.ReadAll() // Lee todo a memoria (slice de slices)

	var sumaNotas float64;var count int

	fmt.Println("Alumnos aprobados:")
	// Saltamos la cabecera (i=1)
	for i := 1; i < len(registros); i++ {
		fila := registros[i] // [Nombre, Edad, Nota]
		nota, _ := strconv.ParseFloat(fila[2], 64)
		
		if nota >= 5.0 {
			fmt.Printf("- %s (%.1f)\n", fila[0], nota)
		}
		sumaNotas += nota
		count++
	}
	fmt.Printf("\nPromedio de la clase: %.2f\n", sumaNotas/float64(count))
	// Limpieza
	os.Remove(archivoNombre)
}
