package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

// Dibuja las pistas de carrera
func DrawWay(screen *ebiten.Image) {
	wayColor := color.RGBA{178, 34, 34, 255}

	for x := 50; x < 250; x++ {
		for y := 25; y < 50; y++{
			screen.Set(x, y, wayColor)
		}
	}

	for x := 50; x < 250; x++ {
		for y := 75; y < 100; y++{
			screen.Set(x, y, wayColor)
		}
	}

	for x := 50; x < 250; x++ {
		for y := 125; y < 150; y++{
			screen.Set(x, y, wayColor)
		}
	}

	for x := 50; x < 250; x++ {
		for y := 175; y < 200; y++{
			screen.Set(x, y, wayColor)
		}
	}

	for x := 50; x < 250; x++ {
		for y := 225; y < 250; y++{
			screen.Set(x, y, wayColor)
		}
	}
}