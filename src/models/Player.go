package models

import "github.com/hajimehoshi/ebiten/v2"

// Genera entidad de corredor
type Player struct {
	Id      int
	X       float64
	Y       float64
	Sprites []*ebiten.Image
}