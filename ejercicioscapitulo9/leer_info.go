//fichero leer_info.go
package main

import (
	"fmt"
	"image"
	_ "image/jpeg" // Importar codecs para que Go sepa leer JPG
	_ "image/png"  // Importar codecs para que Go sepa leer PNG
	"os"
)

func main() {
	// Asegúrate de tener una imagen 'input.jpg' o usa 'patron.png' del ej anterior
	f, err := os.Open("patron.png")
	if err != nil {
		fmt.Println("No encuentro la imagen:", err)
		return
	}
	defer f.Close()

	// DecodeConfig lee solo la cabecera (muy rápido)
	config, format, err := image.DecodeConfig(f)
	if err != nil {
		fmt.Println("Error decodificando:", err)
		return
	}

	fmt.Printf("Formato detectado: %s\n", format)
	fmt.Printf("Dimensiones: %d x %d píxeles\n", config.Width, config.Height)
}
