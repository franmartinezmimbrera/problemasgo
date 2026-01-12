// fichero saludo.go
// Este programa hace un saludo personalizado
package main
import (
	"bufio"
	"fmt"
	"os"
)

func main() {

    fmt.Println("¡Hola! ¿Cómo te llamas?")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() 
	nombre := scanner.Text()
	fmt.Printf("¿Qué tal estás %s?\n", nombre)
}