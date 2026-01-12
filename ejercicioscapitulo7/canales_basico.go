//fichero canales_basico.go
package main

import "fmt"

func main() {
	// Creamos un canal de strings
	mensajes := make(chan string)

	// Goroutine anónima que envía datos
	go func() {
		mensajes <- "Ping" // Enviar al canal (bloqueante si no hay quien reciba)
		mensajes <- "Pong"
		close(mensajes)    // Cerramos el canal para indicar que no hay más datos
	}()

	// Recibir datos del canal
	// El bucle 'range' itera hasta que el canal se cierra
	for msg := range mensajes {
		fmt.Println("Recibido:", msg)
	}
}