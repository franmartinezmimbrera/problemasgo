//fichero select_timeout.go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// simularServidor responde después de un tiempo aleatorio
func simularServidor(nombre string, ch chan string) {
	delay := time.Duration(rand.Intn(1000)) * time.Millisecond
	time.Sleep(delay)
	ch <- fmt.Sprintf("%s respondió en %v", nombre, delay)
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go simularServidor("Servidor A", c1)
	go simularServidor("Servidor B", c2)

	// select espera al PRIMERO que esté listo
	select {
	case res := <-c1:
		fmt.Println("Ganó:", res)
	case res := <-c2:
		fmt.Println("Ganó:", res)
	case <-time.After(500 * time.Millisecond):
		// Si nadie responde en 500ms, se ejecuta este caso
		fmt.Println("Error: Timeout. Los servidores son demasiado lentos.")
	}
}