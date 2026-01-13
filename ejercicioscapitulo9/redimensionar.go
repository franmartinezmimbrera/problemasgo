//fichero redimensionar.go
package main

import (
	"image"
	"image/png"
	"os"
)

// redimensionar reduce a la mitad (algoritmo simple: saltar píxeles)
func redimensionar(src image.Image) image.Image {
	bounds := src.Bounds()
	w, h := bounds.Dx(), bounds.Dy()
	
	newW, newH := w/2, h/2
	dst := image.NewRGBA(image.Rect(0, 0, newW, newH))

	for y := 0; y < newH; y++ {
		for x := 0; x < newW; x++ {
			// Tomamos el píxel de la coordenada doble (2x, 2y)
			dst.Set(x, y, src.At(x*2, y*2))
		}
	}
	return dst
}

func main() {
	f, _ := os.Open("patron.png")
	src, _, _ := image.Decode(f)
	f.Close()

	peque := redimensionar(src)

	out, _ := os.Create("salida_peque.png")
	png.Encode(out, peque)
}
