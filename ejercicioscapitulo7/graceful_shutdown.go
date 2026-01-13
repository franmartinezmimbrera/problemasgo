//fichero graceful_shutdown.go
package main
import ("context";"fmt";"net/http";"os";"os/signal";"time")
func main() {
	srv := &http.Server{Addr: ":8880"}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second) // Simula una petición lenta
		fmt.Fprintln(w, "Hola Mundo")
	})
	// Lanzamos el servidor en una goroutine para que no bloquee
	go func() {
		fmt.Println("Servidor iniciado en puerto 8880 (Ctrl+C para parar)")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error: %v\n", err)
		}
	}()
	// Canal para escuchar señales del sistema operativo
	quit := make(chan os.Signal, 1)
	// Notificarnos si llega SIGINT (Ctrl+C) o SIGTERM (Kubernetes stop)
	signal.Notify(quit, os.Interrupt)
	// Bloqueamos aquí hasta recibir la señal
	<-quit
	fmt.Println("\nApagando servidor...")
	// Damos 5 segundos de margen para terminar peticiones vivas
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Println("Forzando cierre:", err)
	}
	fmt.Println("Servidor apagado correctamente.")
}
