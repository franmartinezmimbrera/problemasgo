//fichero histograma.go
package main

import ("fmt";"image";_ "image/png";"os")

func main() {
	f, _ := os.Open("patron.png")
	img, _, _ := image.Decode(f)
	f.Close()

	// Histograma de Luminosidad (0-255)
	var histograma [256]int
	bounds := img.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			// Convertir a 8 bits y calcular luminosidad simple
			val := (uint8(r>>8) + uint8(g>>8) + uint8(b>>8)) / 3
			histograma[val]++
		}
	}

	fmt.Println("Distribución de luminosidad (resumen):")
	fmt.Printf("Oscuros (0-50): %d px\n", sum(histograma[0:50]))
	fmt.Printf("Medios (51-200): %d px\n", sum(histograma[51:200]))
	fmt.Printf("Claros (201-255): %d px\n", sum(histograma[201:255]))
}

func sum(arr []int) int {
	s := 0
	for _, v := range arr { s += v }
	return s
}
