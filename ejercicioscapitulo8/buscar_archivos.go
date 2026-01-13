//fichero buscar_archivos.go
package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	root := "." // Directorio actual
	if len(os.Args) > 1 {
		root = os.Args[1]
	}

	fmt.Printf("Buscando ficheros .go en '%s'...\n", root)
	// WalkDir recorre recursivamente el árbol
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}	
		// Si no es directorio y acaba en .go
		if !d.IsDir() && filepath.Ext(path) == ".go" {
			info, _ := d.Info()
			fmt.Printf("[GO] %s (%d bytes)\n", path, info.Size())
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error:", err)
	}
}
