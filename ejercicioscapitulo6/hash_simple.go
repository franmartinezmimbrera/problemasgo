// hash_simple.go
package main
import ("bufio";"fmt";"os";"strings")
// hashSimple implementa un checksum básico.
// Suma los valores de los bytes multiplicados por su posición.
func hashSimple(s string) uint64 {
	var hash uint64 = 0
	for i, r := range s {
		// Multiplicamos por (i+1) para que "AB" sea diferente de "BA"
		// Usamos bitwise XOR y desplazamiento para mezclar bits
		hash += uint64(r) * uint64(i+1)
		hash = (hash << 5) ^ hash 
	}
	return hash
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Introduce texto para hashear: ")
	texto1, _ := reader.ReadString('\n')
	texto1 = strings.TrimSpace(texto1)
	h1 := hashSimple(texto1)
	fmt.Printf("Hash 1: %d\n", h1)
	fmt.Print("Introduce texto ligeramente diferente: ")
	texto2, _ := reader.ReadString('\n')
	texto2 = strings.TrimSpace(texto2)
	h2 := hashSimple(texto2)
	fmt.Printf("Hash 2: %d\n", h2)
	if h1 != h2 {
		fmt.Println("¡Integridad verificada! Los hashes son distintos.")
	} else {
		fmt.Println("Colisión o textos idénticos.")
	}
} 
