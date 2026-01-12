// fichero xor_cifra_descifra.go
package main
import ("bufio";"fmt";"os";"strings")
// cifrarDescifrarXOR aplica la operación XOR byte a byte.
// Al ser XOR una operación involutiva ((A ^ B) ^ B = A), la misma función sirve para cifrar y descifrar.
func cifrarDescifrarXOR(input, clave string) string {
	inputBytes := []byte(input)
	claveBytes := []byte(clave)
	output := make([]byte, len(inputBytes))
	for i := 0; i < len(inputBytes); i++ {
		output[i] = inputBytes[i] ^ claveBytes[i%len(claveBytes)]
	}
	return string(output)
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Introduce el mensaje: ")
	mensaje, _ := reader.ReadString('\n')
	mensaje = strings.TrimSpace(mensaje)
	fmt.Print("Introduce la clave: ")
	clave, _ := reader.ReadString('\n')
	clave = strings.TrimSpace(clave)
	resultado := cifrarDescifrarXOR(mensaje, clave)
	fmt.Println("\nResultado")
	fmt.Printf("Hex:   %x\n", resultado)
	recuperado := cifrarDescifrarXOR(resultado, clave)
	fmt.Printf("\nVerificación (Descifrando de nuevo): %s\n", recuperado)
}