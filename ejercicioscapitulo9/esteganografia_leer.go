//fichero esteganografia_leer.go
package main
import ("fmt";"image";_ "image/png";"os")
func main() {
	f, _ := os.Open("secreto.png")
	if f == nil {
		fmt.Println("Ejecuta primero el 9.9 para generar 'secreto.png'")
		return
	}
	img, _, _ := image.Decode(f)
	f.Close()
	bounds := img.Bounds();var bits []int
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			_, _, b, _ := img.At(x, y).RGBA()
			b8 := uint8(b >> 8)	
			// Extraemos último bit
			bits = append(bits, int(b8 & 1))
		}
	}
	var mensaje string
	for i := 0; i < len(bits)-8; i += 8 {
		var charVal int
		for bit := 0; bit < 8; bit++ {
			if bits[i+bit] == 1 {
				charVal |= (1 << bit)
			}
		}
		if charVal == 0 {
			break
		}
		mensaje += string(rune(charVal))
	}
	fmt.Println("Mensaje recuperado:")
	fmt.Println(mensaje)
}
