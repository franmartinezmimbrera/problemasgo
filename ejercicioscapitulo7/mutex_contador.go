//fichero mutex_contador.go
package main
import ("fmt";"sync")
// SafeCounter es seguro para usar concurrentemente
type SafeCounter struct {
	mu sync.Mutex ;v  map[string]int
}
// Inc incrementa el contador para una clave dada
func (c *SafeCounter) Inc(key string) {
	// Bloqueamos el acceso para que solo UNA goroutine escriba a la vez
	c.mu.Lock()
	// Nos aseguramos de desbloquear al terminar
	defer c.mu.Unlock()
	c.v[key]++
}
// Value devuelve el valor actual del contador
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}
func main() {
	c := SafeCounter{v: make(map[string]int)}
	var wg sync.WaitGroup
	// 1000 goroutines intentando incrementar el mismo contador a la vez
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {c.Inc("visitas");wg.Done()
	    }()
	}
	wg.Wait()
	fmt.Println("Valor final (debería ser 1000):", c.Value("visitas"))
}
