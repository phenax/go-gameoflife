package gol

//
// Frame - The properties of a frame in the game
//
type Frame struct {

	// Points
	// Grid of boolean values(alive or dead)
	Points [][]bool
}

//
// NewEmptyFrame - Create an empty frame(all dead)
//
// params
// -- {int} rows  Number of rows
// -- {int} cols  Number of columns
//
// returns
// -- {*Frame}
//
func NewEmptyFrame(rows, cols int) *Frame {

	frame := &Frame{
		Points: make([][]bool, rows, rows),
	}

	for i := range frame.Points {
		frame.Points[i] = make([]bool, cols, cols)
	}

	return frame
}

//
// NewBlinkerFrame - Initial state for the blinker life
//
// params
// -- {int} rows
// -- {int} cols
// -- {int} x  The start position
// -- {int} y  The end position
//
// returns
// -- {*Frame}
//
func NewBlinkerFrame(rows, cols, x, y int) *Frame {

	frame := NewEmptyFrame(rows, cols)

	frame.SetAlive(x, y)
	frame.SetAlive(x, y+1)
	frame.SetAlive(x, y+2)

	return frame
}

//
// NewGliderFrame - Initial state for the glider life
//
// params
// -- {int} rows
// -- {int} cols
// -- {int} x
// -- {int} y
//
// returns
// -- {*Frame}
//
func NewGliderFrame(rows, cols, x, y int) *Frame {

	frame := NewEmptyFrame(rows, cols)

	frame.SetAlive(x, y)
	frame.SetAlive(x+1, y+1)
	frame.SetAlive(x+1, y+2)
	frame.SetAlive(x+2, y)
	frame.SetAlive(x+2, y+1)

	return frame
}

//
// NewSpaceshipFrame - Initial state for the spaceship life
//
// params
// -- {int} rows
// -- {int} cols
// -- {int} x
// -- {int} y
//
// returns
// -- {*Frame}
//
func NewSpaceshipFrame(rows, cols, x, y int) *Frame {

	frame := NewEmptyFrame(rows, cols)

	frame.SetAlive(x, y)
	frame.SetAlive(x, y+2)
	frame.SetAlive(x+1, y+3)
	frame.SetAlive(x+2, y+3)
	frame.SetAlive(x+3, y+3)
	frame.SetAlive(x+4, y+3)
	frame.SetAlive(x+4, y+2)
	frame.SetAlive(x+4, y+1)
	frame.SetAlive(x+3, y)

	return frame
}

//
// NewPulsarFrame - Initial state for the pulsar life
//
// params
// -- {int} rows
// -- {int} cols
// -- {int} x
// -- {int} y
//
// returns
// -- {*Frame}
//
func NewPulsarFrame(rows, cols, x, y int) *Frame {

	frame := NewEmptyFrame(rows, cols)

	height := 7

	frame.SetAlive(x+2, y)
	frame.SetAlive(x+3, y)
	frame.SetAlive(x+4, y)
	frame.SetAlive(x+8, y)
	frame.SetAlive(x+9, y)
	frame.SetAlive(x+10, y)

	for j := 0; j < 3; j++ {
		frame.SetAlive(x, y+2+j)
		frame.SetAlive(x+5, y+2+j)
		frame.SetAlive(x+7, y+2+j)
		frame.SetAlive(x+12, y+2+j)
	}

	frame.SetAlive(x+2, y+5)
	frame.SetAlive(x+3, y+5)
	frame.SetAlive(x+4, y+5)
	frame.SetAlive(x+8, y+5)
	frame.SetAlive(x+9, y+5)
	frame.SetAlive(x+10, y+5)

	frame.SetAlive(x+2, y+height)
	frame.SetAlive(x+3, y+height)
	frame.SetAlive(x+4, y+height)
	frame.SetAlive(x+8, y+height)
	frame.SetAlive(x+9, y+height)
	frame.SetAlive(x+10, y+height)

	for j := 0; j < 3; j++ {
		frame.SetAlive(x, y+1+j+height)
		frame.SetAlive(x+5, y+1+j+height)
		frame.SetAlive(x+7, y+1+j+height)
		frame.SetAlive(x+12, y+1+j+height)
	}

	frame.SetAlive(x+2, y+5+height)
	frame.SetAlive(x+3, y+5+height)
	frame.SetAlive(x+4, y+5+height)
	frame.SetAlive(x+8, y+5+height)
	frame.SetAlive(x+9, y+5+height)
	frame.SetAlive(x+10, y+5+height)

	return frame
}

//
// SetState - Set the state of a cell in frame frame
//
// params
// -- {int} x
// -- {int} y
// -- {bool} state New state
//
func (frame *Frame) SetState(x, y int, state bool) {

	if y < len(frame.Points) && y >= 0 {
		if x < len(frame.Points[y]) && x >= 0 {
			frame.Points[y][x] = state
		}
	}
}

// SetAlive - Set the state of a cell to alive
func (frame *Frame) SetAlive(x, y int) {
	frame.SetState(x, y, true)
}

// SetDead - Set the state of a cell to dead
func (frame *Frame) SetDead(x, y int) {
	frame.SetState(x, y, false)
}

//
// GetAliveNeighbourCount - Get the number of alive neighbors for a cell
//
// params
// -- {int} x
// -- {int} y
//
// returns
// -- {int} Neighbour count
//
func (frame *Frame) GetAliveNeighbourCount(x, y int) int {

	aliveCount := 0

	for i := -1; i <= 1; i++ {

		if y+i < len(frame.Points) && y+i >= 0 {

			for j := -1; j <= 1; j++ {

				if i != 0 || j != 0 {

					row := frame.Points[y+i]

					if x+j < len(row) && x+j >= 0 {

						if row[x+j] {
							aliveCount++
						}
					}
				}
			}
		}
	}

	return aliveCount
}

// String - Convert the frame to a string
func (frame *Frame) String() string {

	str := ""

	frame.ForEachRow(func(row []bool, _ int) {

		for _, point := range row {

			str += " "

			if point {
				str += "#"
			} else {
				str += "-"
			}
		}

		str += "\n"
	})

	return str
}

//
// ForEachRow - Iterate through all rows in a frame
//
func (frame *Frame) ForEachRow(callback func([]bool, int)) {

	for i, row := range frame.Points {
		callback(row, i)
	}
}

//
// ForEach - Iterate through all cells in a frame
//
func (frame *Frame) ForEach(callback func(bool, int, int) bool) {

	frame.ForEachRow(func(row []bool, y int) {

		for x, point := range row {

			continueLoop := callback(point, x, y)

			if !continueLoop {
				return
			}
		}
	})
}
