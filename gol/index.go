package gol

import (
	"fmt"
	"time"
)

type GameOfLife struct {
	Rows         int
	Columns      int
	CurrentFrame *Frame
}

func NewGameOfLife(rows, cols int) *GameOfLife {

	game := &GameOfLife{
		Rows:    rows,
		Columns: cols,
	}

	return game
}

func (this *GameOfLife) LoadFrame(frame *Frame) {
	this.CurrentFrame = frame
}

func (this *GameOfLife) StartLoop() {

	go this.StartCalcLoop()

	for {
		fmt.Print("\033[2J")
		fmt.Print(this.CurrentFrame)
		time.Sleep(200 * time.Millisecond)
	}
}

func (this *GameOfLife) StartCalcLoop() {

	for {
		this.CalculateNextFrame()
		time.Sleep(200 * time.Millisecond)
	}
}

func (this *GameOfLife) CalculateNextFrame() {

	newFrame := EmptyFrame(this.Rows, this.Columns)

	this.CurrentFrame.ForEach(func(point bool, x int, y int) bool {

		neighbourCount := this.CurrentFrame.GetAliveNeighbourCount(x, y)

		if point {
			if neighbourCount == 2 || neighbourCount == 3 {
				newFrame.SetAlive(x, y)
			} else {
				newFrame.SetDead(x, y)
			}
		} else {
			if neighbourCount == 3 {
				newFrame.SetAlive(x, y)
			}
		}

		return true
	})

	this.CurrentFrame = newFrame
}
