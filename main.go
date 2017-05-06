package main

import (
	"github.com/phenax/go-gameoflife/gol"
)

const (
	// Number of rows in the grid
	numberOfRows = 40
	// Number of columns in the grid
	numberOfColumns = 40
)

func main() {

	game := gol.NewGameOfLife(numberOfRows, numberOfColumns)
	initFrame := gol.NewPulsarFrame(numberOfRows, numberOfColumns, 10, 10)

	game.LoadFrame(initFrame)

	// fmt.Println(initFrame)

	game.StartLoop()
}
