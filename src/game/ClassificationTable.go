package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

// Función para dibujar la tabla de clasificación
func DrawClassificationTable(screen *ebiten.Image, classification []int){
	bannerColor := color.RGBA{255,255,255,255}

	// Dibuja el contorno de la tabla
	for x := 0; x < 300; x++ {
		for y := 50; y < 250; y++{
			screen.Set(x, y, bannerColor)
		}
	}

	fountColor := color.RGBA{1,1,1,255}

	var colorClassifications []string

	// Itera sobre la tabla de clasificación y asigna el color
	for i := 0; i < len(classification); i++ {
		switch classification[i] {
		case 0:
			colorClassifications = append(colorClassifications, "Rojo")
		case 1:
			colorClassifications = append(colorClassifications, "Naranja")
		case 2:
			colorClassifications = append(colorClassifications, "Verde")
		case 3:
			colorClassifications = append(colorClassifications, "Azúl")
		case 4:
			colorClassifications = append(colorClassifications, "Púrpura")
		}
	}

	// Escribe la tabla de clasificaciones
	text.Draw(screen, "1. " + colorClassifications[0], ApplyFount(20), 10, 80, fountColor)
	text.Draw(screen, "2. " + colorClassifications[1], ApplyFount(20), 10, 120, fountColor)
	text.Draw(screen, "3. " + colorClassifications[2], ApplyFount(20), 10, 160, fountColor)
	text.Draw(screen, "4. " + colorClassifications[3], ApplyFount(20), 10, 200, fountColor)
	text.Draw(screen, "5. " + colorClassifications[4], ApplyFount(20), 10, 240, fountColor)
}