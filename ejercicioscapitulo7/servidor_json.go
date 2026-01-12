//fichero servidor_json.go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)
// Estructura para la respuesta JSON
type Respuesta struct {
	Mensaje string    `json:"mensaje"`
	Hora    time.Time `json:"hora"`
	Status  int       `json:"status"`
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Petición recibida")
	// Preparamos datos
	data := Respuesta{
		Mensaje: "Hola desde Go!",
		Hora:    time.Now(),
		Status:  200,
	}
	// Configuramos cabeceras
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// Codificamos struct a JSON y enviamos
	json.NewEncoder(w).Encode(data)
}
func main() {
	http.HandleFunc("/api/hola", handler)
	fmt.Println("Servidor escuchando en http://localhost:8880/api/hola")
	// Bloquea indefinidamente
	http.ListenAndServe(":8880", nil)
}
