// fichero ejercicioscapitulo9/escala_grises.go
package main

import ("image";"image/jpeg";_ "image/png";"os")

func main() {
	// Abrimos el fichero original (que es un PNG)
	f, err := os.Open("patron.png") 
	if err != nil {
		panic(err) // Si falla (ej. no existe el archivo), paramos
	}
	defer f.Close()
	// Al tener el import _ "image/png", el Decode sabe leer PNGs automáticamente
	src, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	bounds := src.Bounds()
	// Creamos nueva imagen en escala de grises
	// NewGray crea un lienzo donde cada píxel es solo luz (0-255)
	grayImg := image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// Leemos el color original
			originalColor := src.At(x, y)
			grayImg.Set(x, y, originalColor)
		}
	}
	out, _ := os.Create("salida_gris.jpg")
	defer out.Close()
	// Guardamos con calidad por defecto (nil)
	jpeg.Encode(out, grayImg, nil)
}