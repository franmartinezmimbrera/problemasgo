// fichero memostruc.go
// Ejemplo de manejo de memoria para estructuras de datos
package main
import (
	"fmt"
)
// Definimos la estructura.
type Tarea struct {
	ID     int
	Nombre string
}
func main() {
	nTareas := 10

	listaTareas := make([]Tarea, nTareas)
	listaTareas[0].ID = 1
	listaTareas[0].Nombre = "Comprar"
	listaTareas[1] = Tarea{
		ID:     2,
		Nombre: "Estudiar",
	}
	listaTareas[2] = Tarea{ID: 3, Nombre: "Correr"}
	fmt.Println("Lista de Tareas:")
	// range devuelve índice (i) y una COPIA del valor (tarea).
	for i, tarea := range listaTareas {
		if tarea.ID != 0 {
			fmt.Printf(" - Tarea %d: %s\n", tarea.ID, tarea.Nombre)
		} else {
			fmt.Printf(" - Tarea %d: <No asignada>\n", i+1)
		}
	}
	// Liberación Automática por el Garbage Collector 
	fmt.Println("Memoria de estructuras liberada correctamente (GC).")
}