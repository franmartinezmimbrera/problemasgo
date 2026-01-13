//fichero recortar.go
package main

import (
	"image"
	"image/png"
	"os"
)

func main() {
	f, _ := os.Open("patron.png")
	src, _, _ := image.Decode(f)
	defer f.Close()

	// Interfaz SubImage (si la imagen la soporta)
	type SubImager interface {
		SubImage(r image.Rectangle) image.Image
	}

	if img, ok := src.(SubImager); ok {
		// Definimos el rectángulo de recorte (x0, y0, x1, y1)
		rect := image.Rect(50, 50, 150, 150)
		crop := img.SubImage(rect)

		out, _ := os.Create("salida_crop.png")
		png.Encode(out, crop)
	}
}
