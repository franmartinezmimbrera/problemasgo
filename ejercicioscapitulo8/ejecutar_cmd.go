//fichero ejecutar_cmd.go
package main

import (
	"fmt"
	"os/exec"
)

func main() {

	// Vamos a ejecutar un comando 'ping' (o 'ls' / 'dir')
	// Nota: En Windows usar "ping", "-n", "3", "google.com"
	
	cmd := exec.Command("ping", "-c", "3", "google.com")

	fmt.Println("Ejecutando comando externo...")
	
	// Output ejecuta el comando y devuelve la salida estándar capturada
	salida, err := cmd.Output()
	
	if err != nil {
		fmt.Println("El comando falló:", err)
		return
	}

	fmt.Println("Resultado:")
	fmt.Println(string(salida))
}
