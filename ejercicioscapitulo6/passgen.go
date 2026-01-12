package main
import ("crypto/rand";"fmt";"math/big")
const (
	letrasMay = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letrasMin = "abcdefghijklmnopqrstuvwxyz"
	numeros   = "0123456789"
	simbolos  = "!@#$%^&*()-_=+"
	todo      = letrasMay + letrasMin + numeros + simbolos
)
// generarPassword usa crypto/rand para seguridad criptográfica
func generarPassword(longitud int) (string, error) {
	bytes := make([]byte, longitud)
	max := big.NewInt(int64(len(todo)))
	for i := 0; i < longitud; i++ {
		num, err := rand.Int(rand.Reader, max)
		if err != nil {return "", err}
		bytes[i] = todo[num.Int64()]
	}
	return string(bytes), nil
}
func main() {
	longitud := 16
	password, err := generarPassword(longitud)
	if err != nil {
		fmt.Println("Error generando password:", err)
		return
	}
	fmt.Printf("Generador de Contraseñas Seguras (crypto/rand)\n")
	fmt.Printf("Longitud: %d\n", longitud)
	fmt.Printf("Password: %s\n", password)
}