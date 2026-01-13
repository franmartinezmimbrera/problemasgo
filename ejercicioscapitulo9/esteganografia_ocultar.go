//fichero esteganografia_ocultar.go
package main
import ("fmt";"image";"image/color";"image/png";"os")
func main() {
	mensaje := "Este es un mensaje secreto top secret!!"
	// Añadimos un caracter nulo al final para saber dónde parar de leer
	mensaje += string(rune(0)) 
	f, _ := os.Open("patron.png")
	img, _, _ := image.Decode(f)
	f.Close()
	bounds := img.Bounds();nuevo := image.NewRGBA(bounds)
	// Convertimos mensaje a bits
	bits := make([]int, 0)
	for _, char := range mensaje {
		for i := 0; i < 8; i++ {bit := (int(char) >> i) & 1;bits = append(bits, bit)}
	}
	bitIndex := 0
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			c := color.RGBA{uint8(r>>8), uint8(g>>8), uint8(b>>8), uint8(a>>8)}
			if bitIndex < len(bits) {
			// Ponemos a 0 el último bit de Azul y sumamos el bit del mensaje
			// (c.B & 0xFE) borra el último bit
				c.B = (c.B & 0xFE) | uint8(bits[bitIndex])
				bitIndex++
			}
			nuevo.Set(x, y, c)
		}
	}
	out, _ := os.Create("secreto.png")
	png.Encode(out, nuevo)
	fmt.Println("Mensaje oculto en secreto.png")
}
