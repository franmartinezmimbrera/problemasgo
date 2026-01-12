// fichero sumaparesimpares.go
// Calcula la suma de 10 números pasados por teclado distinguiendo entre los pares e impares
package main

import ("bufio";"fmt";"os";"strconv";"strings")
const TamanoArray = 10

func main() {
	numeros := make([]int64, TamanoArray)
	var sumaPares, sumaImpares int64;var err error
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Por favor, introduce %d números enteros:\n", TamanoArray)
	for i := 0; i < TamanoArray; i++ {
		fmt.Printf("Número %d: ", i+1)
		if scanner.Scan() {
			texto := strings.TrimSpace(scanner.Text())
			numeros[i], err = strconv.ParseInt(texto, 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error en la lectura o entrada no válida. Terminando...")
				os.Exit(1)
			}
		} else {
			os.Exit(1)
		}
	}
	for _, numeroActual := range numeros {
		if numeroActual%2 == 0 {
			sumaPares += numeroActual
			fmt.Printf("Añadiendo %d a la suma par.\n", numeroActual)
		} else {
			sumaImpares += numeroActual
			fmt.Printf("Añadiendo %d a la suma impar.\n", numeroActual)
		}
	}
	fmt.Print("Números ingresados: [")
	for i, num := range numeros {
		fmt.Print(num)
		if i < len(numeros)-1 {fmt.Print(", ")}
	}
	fmt.Println("]")
	fmt.Printf("Suma total de los números PARES: %d\n", sumaPares)
	fmt.Printf("Suma total de los números IMPARES: %d\n", sumaImpares)
}