package threads

import (
	"context"
	"sync"
	"time"

	"programa.concurrencia/src/game"
)

// Función que administra fan-out/fan-in
func ExecuteFanInFanOutPattern(game *game.GameManager, ctx context.Context){
	// Crea canal y mutex
	resultsChan := make(chan int)
	mutex := sync.Mutex{}

	// Al terminar, cierra canal
	defer close(resultsChan)

	// Define distancia en pixeles de la carrera
	race_duration := 190

	println("Iniciando carrera")

	n_runners := 5

	// Por cada corredor, crear una goroutine
	for i := 0; i < n_runners; i++ {
		go WorkerFanOutFanIn(game, resultsChan, &mutex, ctx, race_duration, i)
	}

	var classifications []int
	// Espera señales
	for i := 0; i < n_runners; i++ {
		select{
			// Si llega señal del context, termina esta goroutine
			case <- ctx.Done():
				println("Cancelando carrera")
				return
			// Si llega señal del canal, va agregando los jugadores en orden de llegada
			case place := <- resultsChan:
				classifications = append(classifications, place)
		}
		
	}

	// Activa la tabla de clasificación en pantalla
	game.SetClassification(classifications, true)

	// Espera 5 segundos
	time.Sleep(10 * time.Second)

	// Termina juego
	game.FinishGame()
}