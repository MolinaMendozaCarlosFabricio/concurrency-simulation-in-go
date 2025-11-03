package threads

import (
	"context"
	"math/rand"
	"sync"
	"time"

	"programa.concurrencia/src/game"
)

// Función del worker select-multiplexing
func ParticipantRun(
	game *game.GameManager,
	finishChan chan struct{},
	mutex *sync.Mutex,
	ctx context.Context,
	duration int,
	runner_id int,
) {
	// Cierra su propio canal al concluir
	defer close(finishChan)
	for i := 0; i <= duration; i++ {
		select {
			// Si recibe señal del context, cierra esta goroutine
			case<-ctx.Done():
				println("Worker cerrado")
				return
			default:
				// Aplica mutex para modificar atributo de corredores
				// en el struct del juego
				mutex.Lock()
				// Desplaza el corredor en un pixel
				game.UpdatePlayerCoordinates(runner_id)
				mutex.Unlock()
				// Espera un tiempo aleatorio
				random_duration := 100 * rand.Intn(3)
				time.Sleep(time.Duration(random_duration) * time.Millisecond)
		}
	}

	// Envía señal a su propio canal
	finishChan <- struct{}{}
}