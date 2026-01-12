//fichero contexto_cancel.go
package main

import (
	"context"
	"fmt"
	"time"
)

func operacionLenta(ctx context.Context) {
	// select permite escuchar al canal Done() del contexto
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("Operación completada con éxito")
	case <-ctx.Done():
		// Si el contexto se cancela o expira, entramos aquí inmediatamente
		fmt.Println("Operación cancelada:", ctx.Err())
	}
}
func main() {
	// Creamos un contexto que expira automáticamente en 1 segundo
	// (Timeout)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	
	// Es buena práctica llamar a cancel al salir, aunque haya expirado
	defer cancel()

	fmt.Println("Iniciando operación que tarda 2 segundos...")
	// Pasamos el contexto a la función
	operacionLenta(ctx)
	
	fmt.Println("Programa terminado")
}
