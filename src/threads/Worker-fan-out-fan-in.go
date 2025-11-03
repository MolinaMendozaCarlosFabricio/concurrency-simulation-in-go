package threads

import (
	"context"
	"math/rand"
	"sync"
	"time"

	"programa.concurrencia/src/game"
)

// Función de worker fan-out/fan-int
func WorkerFanOutFanIn(
	game *game.GameManager,
	finishChan chan int,
	mutex *sync.Mutex,
	ctx context.Context,
	duration int,
	runner_id int,
){
	for i := 0; i <= duration; i++ {
		select {
			// Si llega señal del context, termina esta goroutine
			case <- ctx.Done():
				println("Finalizando worker")
				return
			default:
				// Aplica mutex para modificar atributo de corredores
				// en el struct del juego
				mutex.Lock()
				// Avanza un pixel más
				game.UpdatePlayerCoordinates(runner_id)
				mutex.Unlock()
				// Espera un tiempo aleatorio
				random_duration := 100 * rand.Intn(3)
				time.Sleep(time.Duration(random_duration) * time.Millisecond)
		}
	}

	// Envía ID del corredor por el canal
	finishChan <- runner_id
}