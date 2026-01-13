//fichero generar_patron.go
package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	width, height := 256, 256
	// Creamos un lienzo RGBA en blanco
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			// Generamos un degradado:
			// Rojo aumenta con X, Verde aumenta con Y, Azul fijo
			c := color.RGBA{
				R: uint8(x), 
				G: uint8(y), 
				B: 100, 
				A: 255, // Alfa opaco
			}
			img.Set(x, y, c)
		}
	}

	f, _ := os.Create("patron.png")
	defer f.Close()
	png.Encode(f, img)
}
