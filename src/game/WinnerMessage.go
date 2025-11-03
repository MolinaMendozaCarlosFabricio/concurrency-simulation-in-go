package game

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

func DrawWinnerMessage(screen *ebiten.Image, message string) {
	bannerColor := color.RGBA{255,255,255,255}

	// Dibuja el contorno del mensaje
	for x := 0; x < 300; x++ {
		for y := 125; y < 175; y++{
			screen.Set(x, y, bannerColor)
		}
	}

	fountColor := color.RGBA{1,1,1,255}

	// Escribe mensaje
	text.Draw(screen, message, ApplyFount(20), 10, 160, fountColor)
}