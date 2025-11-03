package threads

import (
	"context"
	"sync"
	"time"

	"programa.concurrencia/src/game"
)

// Función para ejecutra select-multiplexing
func ExecuteSelectMultiplexingPattern(game *game.GameManager, ctx context.Context) {
	// Inicializa canales
	runner1_chan := make(chan struct{})
	runner2_chan := make(chan struct{})
	runner3_chan := make(chan struct{})
	runner4_chan := make(chan struct{})
	runner5_chan := make(chan struct{})

	// Crea mutex
	mutex := sync.Mutex{}

	// Define distancia en pixeles de la carrera
	raceDuration := 190

	println("Iniciando carrera")

	// Genera goroutines con base a los canales
	go ParticipantRun(game, runner1_chan, &mutex, ctx, raceDuration, 0)
	go ParticipantRun(game, runner2_chan, &mutex, ctx, raceDuration, 1)
	go ParticipantRun(game, runner3_chan, &mutex, ctx, raceDuration, 2)
	go ParticipantRun(game, runner4_chan, &mutex, ctx, raceDuration, 3)
	go ParticipantRun(game, runner5_chan, &mutex, ctx, raceDuration, 4)

	var winner string

	// Almacena ganador, con base al canal del quien reciba señal primero
	select {
		case <-runner1_chan:
			println("Ganó corredor rojo")
			winner = "Jugador rojo gana"
		case <-runner2_chan:
			println("Ganó corredor naranja")
			winner = "Jugador naranja gana"
		case <-runner3_chan:
			println("Ganó corredor verde")
			winner = "Jugador verde gana"
		case <-runner4_chan:
			println("Ganó corredor azúl")
			winner = "Jugador azúl gana"
		case <-runner5_chan:
			println("Ganó corredor púrpura")
			winner = "Jugador púrpura gana"
		// Si es señal del context, acaba esta goroutine
		case <-ctx.Done():
			println("Cancelando carrera")
			return
	}

	// Imprime ganador en pantalla
	game.SetWinner(winner, true)

	// Espera 5 segundos
	time.Sleep(10 * time.Second)

	// Cierra juego
	game.FinishGame()
}