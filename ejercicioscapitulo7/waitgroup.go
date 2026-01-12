//fichero waitgroup.go
package main

import (
	"fmt"
	"sync"
	"time"
)
// worker simula una tarea que tarda un tiempo
func worker(id int, wg *sync.WaitGroup) {
	// defer asegura que Done() se ejecute al salir de la función
	defer wg.Done()

	fmt.Printf("Worker %d iniciando...\n", id)
	time.Sleep(time.Second) // Simular trabajo pesado
	fmt.Printf("Worker %d completado.\n", id)
}

func main() {
	var wg sync.WaitGroup

	// Vamos a lanzar 3 goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1) // Incrementamos el contador
		go worker(i, &wg)
	}

	fmt.Println("Esperando a que terminen los workers...")
	
	// Wait() bloquea hasta que el contador del WaitGroup llegue a 0
	wg.Wait()
	
	fmt.Println("Todos los procesos han terminado.")
}