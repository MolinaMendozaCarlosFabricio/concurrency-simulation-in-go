package loadsprites

import (
	"image"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"programa.concurrencia/src/models"
)

// Función para cargar sprites de un corredor
func LoadSprites(id int, x, y float64, frame1 string, frame2 string) models.Player {
	imgs := []string{frame1, frame2}
	// Arreglo de frames
	var frames []*ebiten.Image

	for _, path := range imgs {
		// Obtiene los frames del sprite
		f, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		img, _, err := image.Decode(f)
		if err != nil {
			log.Fatal(err)
		}

		frames = append(frames, ebiten.NewImageFromImage(img))
	}

	// Genera una instancia de un persona y lo retorna
	return models.Player{Id: id, X: x, Y: y, Sprites: frames}
}