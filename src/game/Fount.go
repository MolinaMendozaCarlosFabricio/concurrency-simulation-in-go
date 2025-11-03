package game

import (
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

// Función para aplicar la fuente
func ApplyFount(size float64) font.Face {
    // busca el archivo de fuente
	fontBytes, _ := os.ReadFile("assets/04B_03__.TTF")
    // Parsea la fuente
    ttf, _ := opentype.Parse(fontBytes)
    // Fuente aplicable, a partir de un tamaño dado
    face, _ := opentype.NewFace(ttf, &opentype.FaceOptions{
        Size: size,
        DPI:  72,
    })
    return face
}