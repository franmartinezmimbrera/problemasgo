//fichero filtro_negativo.go
package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)
func main() {
	f, _ := os.Open("patron.png")
	src, _, _ := image.Decode(f)
	defer f.Close()

	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := src.At(x, y).RGBA()
			// Nota: .RGBA() devuelve valores de 16-bits (0-65535)
			// Convertimos a 8-bits
			r8, g8, b8 := uint8(r>>8), uint8(g>>8), uint8(b>>8)
			// Invertimos el color (255 - valor)
			negColor := color.RGBA{
				R: 255 - r8,
				G: 255 - g8,
				B: 255 - b8,
				A: uint8(a >> 8),
			}
			dst.Set(x, y, negColor)
		}
	}
	out, _ := os.Create("salida_negativo.png")
	png.Encode(out, dst)
}
