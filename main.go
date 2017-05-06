package main

import (
	"github.com/phenax/go-gameoflife/gol"
)

const (
	numberOfRows    = 20
	numberOfColumns = 20
)

func main() {

	game := gol.NewGameOfLife(numberOfRows, numberOfColumns)
	initFrame := gol.EmptyFrame(numberOfRows, numberOfColumns)

	createInitialFrame(initFrame)

	game.LoadFrame(initFrame)

	game.StartLoop()
}

func createInitialFrame(frame *gol.Frame) {
	frame.SetAlive(5, 5)
	frame.SetAlive(5, 6)
	frame.SetAlive(5, 7)
}
