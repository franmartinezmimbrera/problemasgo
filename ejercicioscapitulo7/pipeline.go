//fichero pipeline.go
package main

import "fmt"

// Genera números
func generar(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {out <- n}
		close(out)
	}()
	return out
}
// Calcula el cuadrado
func cuadrado(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {out <- n * n}
		close(out)
	}()
	return out
}

func main() {

	// entrada -> generar -> cuadrado -> salida
	c1 := generar(2, 3, 4, 5)
	c2 := cuadrado(c1)
	// Consumimos el resultado final
	for result := range c2 {
		fmt.Println(result)
	}
}
