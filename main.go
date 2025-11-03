package main

import (
	"context"
	"errors"
	"flag"

	"github.com/hajimehoshi/ebiten/v2"
	"programa.concurrencia/src/game"
	loadsprites "programa.concurrencia/src/loadSprites"
	"programa.concurrencia/src/models"
	"programa.concurrencia/src/threads"
)

func main() {
	// Declara contexto
	ctx, cancel := context.WithCancel(context.Background())
	// flag para ejecutar uno de los patrones de concurrencia
	pattern := flag.String("pattern", "select-multiplexing", "Patrón de concurrencia a ejecutar")

	// Configuración de pantalla
	ebiten.SetWindowSize(800, 500)
	ebiten.SetWindowTitle("Simulación de Concurrencia")	

	var players []models.Player

	// Carga los sprites de los jugadores
	players =  append(
		players, 
		loadsprites.LoadSprites(1, 65, 25, "assets/runner_1_1.png", "assets/runner_1_2.png"),
		loadsprites.LoadSprites(2, 65, 75, "assets/runner_2_1.png", "assets/runner_2_2.png"),
		loadsprites.LoadSprites(3, 65, 125, "assets/runner_3_1.png", "assets/runner_3_2.png"),
		loadsprites.LoadSprites(4, 65, 175, "assets/runner_4_1.png", "assets/runner_4_2.png"),
		loadsprites.LoadSprites(5, 65, 225, "assets/runner_5_1.png", "assets/runner_5_2.png"),
	)

	flag.Parse()
	println(pattern)

	// Decide qué patrón usar
	var use_pattern string
	if *pattern == "fan-out/fan-in" {
		use_pattern = "Fan-Out/Fan-In"
	} else {
		use_pattern = "Select-Multiplexing"
	}

	// Define un estatus de salida segura para el juego
	exitStatus := errors.New("exit game")

	// Genera juego
	game := game.NewGame(players, use_pattern, exitStatus, ctx, cancel)

	if use_pattern == "Fan-Out/Fan-In" {
		// Ejecuta patrón Fan-out/fan-in
		go threads.ExecuteFanInFanOutPattern(game, ctx)
	} else {
		// Ejecuta patrón select-multiplexing
		go threads.ExecuteSelectMultiplexingPattern(game, ctx)
	}

	// Ejecuta juego
	if err := ebiten.RunGame(game); err != nil && err != exitStatus {
		panic(err)
	}	
}