//fichero worker_pool.go
package main
import (
	"fmt"
	"time"
)
// worker es el "empleado" que procesa tareas
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d comenzó tarea %d\n", id, j)
		time.Sleep(500 * time.Millisecond) // Simula trabajo costoso
		fmt.Printf("Worker %d terminó tarea %d\n", id, j)
		results <- j * 2
	}
}
func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	// Lanzamos 3 workers (goroutines fijas)
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	// Enviamos 5 tareas
	for j := 1; j <= numJobs; j++ {jobs <- j}
	close(jobs) // Cerramos canal para que los workers sepan que no hay más trabajo
	// Recogemos resultados
	for a := 1; a <= numJobs; a++ {<-results}
	fmt.Println("Todas las tareas procesadas.")
}
