//fichero analizar_disco.go
package main

import (
	"fmt"
	"os"
)

func main() {
	dir := "."
	entries, _ := os.ReadDir(dir)

	var total int64
	fmt.Printf("Analizando '%s'...\n", dir)

	for _, e := range entries {
		info, _ := e.Info()
		size := info.Size()
		total += size
		
		if size > 1000 { // Destacar ficheros > 1KB
			fmt.Printf("! GRANDE: %-20s %d bytes\n", e.Name(), size)
		} else {
			fmt.Printf("  Pequeño: %-20s %d bytes\n", e.Name(), size)
		}
	}
	fmt.Printf("\nTotal ocupado: %d bytes\n", total)
}
