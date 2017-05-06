package main

import (
	"github.com/phenax/go-gameoflife/gol"
)

const (
	numberOfRows    = 40
	numberOfColumns = 40
)

func main() {

	game := gol.NewGameOfLife(numberOfRows, numberOfColumns)
	initFrame := gol.NewPulsarFrame(numberOfRows, numberOfColumns, 10, 10)

	game.LoadFrame(initFrame)

	// fmt.Println(initFrame)

	game.StartLoop()
}
