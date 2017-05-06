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
	initFrame := gol.BlinkerFrame(numberOfRows, numberOfColumns, 6, 6)

	game.LoadFrame(initFrame)

	game.StartLoop()
}
