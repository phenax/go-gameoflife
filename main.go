package main

import (
	"github.com/phenax/go-gameoflife/gol"
)

const (
	Rows    = 50
	Columns = 50
)

func main() {

	game := gol.NewGameOfLife(Rows, Columns)

	initFrame := gol.NewFrame(Rows, Columns)
	createInitialFrame(initFrame)

	game.LoadFrame(initFrame)

	game.StartLoop()
}

func createInitialFrame(frame *gol.Frame) {
	frame.SetAlive(5, 5)
	frame.SetAlive(6, 6)
	frame.SetAlive(7, 7)
}
