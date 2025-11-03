package game

import (
	"context"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"programa.concurrencia/src/models"
)

type GameState int

// Dimenciones en pixeles
const (
	screenWidth  = 300
	screenHeight = 300
)

// Manejo de los frames de los jugadores
var (
	frameCounter = 0
    frameMovemente = 1
)

// Atributos del juego
type GameManager struct{
	players []models.Player //Instancias de jugadores
	frame_player int // Frame de los jugadores
	showWinner bool // Muestra ganador, si se usa select-multiplexing
	winnerMessage string // Mensaje de ganador
	showClassification bool // Muestra tabla de clasificación, si se usa fan-out/fan-in
	classification []int // Tabla de clasificación
	pattern string // Patrón que se está usando
	cancelContext context.CancelFunc // Función para cancelar contexto
	finish bool // Marca fin del programa
	exitStatus error // Estatus de salida
}

// Genera instancia del juego
func NewGame(
	players []models.Player, 
	pattern string, 
	exitStatus error, 
	ctx context.Context,
	cancel context.CancelFunc,
)*GameManager{
	return&GameManager{
		players: players, 
		frame_player: 0,
		showWinner: false,
		winnerMessage: "",
		showClassification: false,
		pattern: pattern,
		finish: false,
		exitStatus: exitStatus,
		cancelContext: cancel,
	}
}

// Define dimenciones del display
func (g *GameManager) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

// Ejecuta lógica del juego por cada tick
func (g *GameManager) Update() error {
	// Al presionar Esc o finalizar desde otra goroutine, termina juego
	if ebiten.IsKeyPressed(ebiten.KeyEscape) || g.finish {
		println("Cancelando contexto y cerrando juego...")
		// Cierra todas las goroutines
		g.cancelContext()
		return g.exitStatus
	}
	return nil
}

// Renderiza elementos
func (g *GameManager) Draw(screen *ebiten.Image) {
	// Efecto de cesped
	screen.Fill(color.RGBA{124, 252, 0, 255})
	// Dibuja pista
	DrawWay(screen)

	// Escribe los índices de los jugadores
	text.Draw(screen, "0", ApplyFount(25), 10, 50, color.RGBA{1,1,1,255})
	text.Draw(screen, "1", ApplyFount(25), 10, 100, color.RGBA{1,1,1,255})
	text.Draw(screen, "2", ApplyFount(25), 10, 150, color.RGBA{1,1,1,255})
	text.Draw(screen, "3", ApplyFount(25), 10, 200, color.RGBA{1,1,1,255})
	text.Draw(screen, "4", ApplyFount(25), 10, 250, color.RGBA{1,1,1,255})

	// Escribe el patrón que se ejecuta actualmente
	if g.pattern == "Select-Multiplexing" {
		ebitenutil.DebugPrint(screen, "Ejecutando carrera usando patrón Select-Multiplexing")
	}
	if g.pattern == "Fan-Out/Fan-In" {
		ebitenutil.DebugPrint(screen, "Ejecutando carrera usando patrón Fan-Out/Fan-In")
	}

	// Define el frame a ejecutar de los corredores
	frameCounter++
	if frameCounter%10 == 0 {
		if g.frame_player == 0 {
			g.frame_player = 1
		} else {
			g.frame_player = 0
		}
	}

	// Actualiza ubicación en pantalla de los sprates de los jugadores
	for i := range g.players {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(-0.1, 0.1)
		op.GeoM.Translate(g.players[i].X, g.players[i].Y)
		screen.DrawImage(g.players[i].Sprites[g.frame_player], op)
	}

	// Muestra mensaje de ganador
	if g.showWinner {
		DrawWinnerMessage(screen, g.winnerMessage)
	}
	// Muestra tabla de clasificación
	if g.showClassification {
		DrawClassificationTable(screen, g.classification)
	}
}

// Asigna ubicación de un corredor, a partir de su índice, o id
func (g *GameManager) UpdatePlayerCoordinates(id int){
	g.players[id].X += 1
}

// Ingresa el jugador ganador, y activa mensaje en pantalla
func (g *GameManager) SetWinner(message string, show bool){
	g.showWinner = true
	g.winnerMessage = message
}

// Ingresa tabla de clasificación, y activa tabla en pantalla
func (g *GameManager) SetClassification(classification []int, show bool){
	g.showClassification = show
	g.classification = classification
}

// Activa el estado de cierre de juego
func (g *GameManager) FinishGame(){
	g.finish = true
}