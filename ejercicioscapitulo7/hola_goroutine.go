// hola_goroutine.go
package main
import (
	"fmt"
	"time"
)
func imprimir(mensaje string) {
	for i := 1; i <= 5; i++ {
		time.Sleep(100 * time.Millisecond) // Simula trabajo
		fmt.Printf("%s: %d\n", mensaje, i)
	}
}
func main() {
	// "go" lanza la función en un hilo aparte
	go imprimir("Goroutine")
	// Ejecutamos otra cosa en el hilo principal
	imprimir("Main")
	time.Sleep(1 * time.Second)
	fmt.Println("Fin del programa")
}