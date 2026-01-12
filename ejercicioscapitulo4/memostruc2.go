// fichero memostruc2.go
// Ejemplo completo: Estructuras en memoria dinámica y persistencia en fichero.
package main

import (
	"encoding/gob"
	"fmt"
	"io"
	"os"
)
type Tarea struct {
	ID     int
	Nombre string
}
const NombreFichero = "registroestructura.gob"

func main() {

	// Creamos un Slice con datos iniciales
	listaOriginal := []Tarea{
		{ID: 1, Nombre: "Comprar"},
		{ID: 2, Nombre: "Estudiar"},
		{ID: 3, Nombre: "Correr"},
	}
	// Creamos el fichero
	fichero, err := os.Create(NombreFichero)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error al crear fichero: %v\n", err)
		os.Exit(1)
	}
	// Creamos el codificador (Encoder)
	encoder := gob.NewEncoder(fichero)

	for _, tarea := range listaOriginal {
		err := encoder.Encode(tarea)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error guardando tarea: %v\n", err)
		}
	}
	
	fichero.Close() 
	fmt.Println("Vector de escritura liberado automáticamente (GC).")
	fmt.Println("Datos guardados correctamente en disco.")

	// Abrimos el fichero para lectura
	ficheroLectura, err := os.Open(NombreFichero)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error al abrir fichero: %v\n", err)
		os.Exit(1)
	}
	defer ficheroLectura.Close()

	decoder := gob.NewDecoder(ficheroLectura)
	var listaLeida []Tarea 
	var numRegistros int

	fmt.Println("\nLista de Tareas leídas del fichero:")
	for {
		var t Tarea
		err := decoder.Decode(&t)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error de lectura: %v\n", err)
			break
		}
		listaLeida = append(listaLeida, t)
		fmt.Printf(" - Tarea %d: %s\n", t.ID, t.Nombre)
		numRegistros++
	}
	fmt.Printf("Se leyeron %d registros.\n", numRegistros)
	fmt.Println("Vector de lectura liberado automáticamente (GC).")
}