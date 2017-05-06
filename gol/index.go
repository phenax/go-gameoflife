package gol

import (
	"fmt"
	"time"
)

const (
	// FrameDelay - Time to wait before rendering/calculating the next frame
	FrameDelay = 200 * time.Millisecond
)

//
// GameOfLife - gol api
//
type GameOfLife struct {

	// Number of rows in the game
	Rows int

	// Number of columns
	Columns int

	// Current rendered frame in the game
	CurrentFrame *Frame
}

//
// NewGameOfLife - Create a new game
//
// params
// -- {int} rows  Number of rows
// -- {int} cols  Number of columns
//
// returns
// -- {*GameOfLife}
//
func NewGameOfLife(rows, cols int) *GameOfLife {

	game := &GameOfLife{
		Rows:    rows,
		Columns: cols,
	}

	return game
}

//
// LoadFrame - Load a new/inital frame into the game
//
// params
// -- {*Frame} frame
//
func (frame *GameOfLife) LoadFrame(frame *Frame) {
	frame.CurrentFrame = frame
}

//
// StartLoop - Start the calculation and render loop
//
func (frame *GameOfLife) StartLoop() {

	// Run the calc loop on a new go routine
	go frame.StartCalcLoop()

	for {
		// Clear the terminal
		fmt.Print("\033[2J")
		// Print the frame
		fmt.Print(frame.CurrentFrame)
		// Sleep for some time
		time.Sleep(FrameDelay)
	}
}

//
// StartCalcLoop - Start the calculation loop
//
func (frame *GameOfLife) StartCalcLoop() {

	for {
		// Calculate the next frame
		frame.CalculateNextFrame()
		// Sleep for some time
		time.Sleep(FrameDelay)
	}
}

//
// CalculateNextFrame - Calculate the next frame in the life
//
func (frame *GameOfLife) CalculateNextFrame() {

	newFrame := NewEmptyFrame(frame.Rows, frame.Columns)

	frame.CurrentFrame.ForEach(func(point bool, x int, y int) bool {

		neighbourCount := frame.CurrentFrame.GetAliveNeighbourCount(x, y)

		// GOL rules
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

	frame.CurrentFrame = newFrame
}
