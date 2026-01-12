// fichero creabinario.go
// Ejemplo de serialización de una estructura en fichero binario
package main

import (
	"encoding/gob"
	"fmt"
	"os"
)
// Definición de la estructura.
// IMPORTANTE: Los campos deben empezar por Mayúscula (Exportados)
// para que el paquete 'gob' pueda leerlos y guardarlos.
type RegistroAlumno struct {
	ID     int
	Nombre string
	Nota   float32
}
const NombreFichero = "registro.gob"
// crearBinario crea el fichero y guarda la estructura serializada
func crearBinario() {
	alumnoOriginal := RegistroAlumno{
		ID:     101,
		Nombre: "Sofia Perez",
		Nota:   9.5,
	}
	fichero, err := os.Create(NombreFichero)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error al crear el fichero: %v\n", err)
		return
	}
	defer fichero.Close()

	// El encoder se encarga de traducir la estructura a bytes binarios
	encoder := gob.NewEncoder(fichero)

	// Go detecta automáticamente que hay un string, calcula su largo,
	// lo guarda y guarda los datos. No hace falta hacerlo manual.
	err = encoder.Encode(alumnoOriginal)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error al guardar los datos: %v\n", err)
		return
	}

	fmt.Println("Fichero binario creado correctamente.")
}
// leerBinario lee del fichero y reconstruye la estructura
func leerBinario() {

	var alumnoLeido RegistroAlumno
	fichero, err := os.Open(NombreFichero)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error al leer el fichero: %v\n", err)
		return
	}
	defer fichero.Close()

	decoder := gob.NewDecoder(fichero)
	// Pasamos un puntero (&) para que Go rellene la variable
	err = decoder.Decode(&alumnoLeido)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error al decodificar los datos: %v\n", err)
		return
	}

	fmt.Printf("Datos cargados del fichero: ID=%d, Nombre=%s, Nota=%.2f\n",
		alumnoLeido.ID, alumnoLeido.Nombre, alumnoLeido.Nota)
}
func main() {
	crearBinario()
	leerBinario()
}