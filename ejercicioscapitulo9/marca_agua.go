//fichero marga_agua.go
package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
)
func main() {
	// Cargar imagen base
	f, _ := os.Open("patron.png")
	base, _, _ := image.Decode(f)
	f.Close()

	b := base.Bounds()
	// Necesitamos una imagen editable (DrawImage)
	lienzo := image.NewRGBA(b)
	// Copiamos la base al lienzo
	draw.Draw(lienzo, b, base, image.Point{}, draw.Src)
	// Crear la marca de agua (un cuadrado rojo semitransparente)
	marca := image.NewRGBA(image.Rect(0, 0, 50, 50))
	rojoTransparente := color.RGBA{255, 0, 0, 128} // Alpha 128 (50%)
	draw.Draw(marca, marca.Bounds(), &image.Uniform{rojoTransparente}, image.Point{}, draw.Src)

	// Dibujar marca sobre el lienzo (offset 20,20)
	offset := image.Pt(20, 20)
	destRect := image.Rect(20, 20, 70, 70) // Dónde va a ir
	// draw.Over superpone mezclando el canal Alpha
	draw.Draw(lienzo, destRect, marca, image.Point{}, draw.Over)
	out, _ := os.Create("salida_watermark.png")
	png.Encode(out, lienzo)
}
