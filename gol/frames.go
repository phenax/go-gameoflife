package gol

type Frame struct {
	Points [][]bool
}

func NewEmptyFrame(rows, cols int) *Frame {

	frame := &Frame{
		Points: make([][]bool, rows, rows),
	}

	for i := range frame.Points {
		frame.Points[i] = make([]bool, cols, cols)
	}

	return frame
}

func NewBlinkerFrame(rows, cols, x, y int) *Frame {

	frame := NewEmptyFrame(rows, cols)

	frame.SetAlive(x, y)
	frame.SetAlive(x, y+1)
	frame.SetAlive(x, y+2)

	return frame
}

func NewGliderFrame(rows, cols, x, y int) *Frame {

	frame := NewEmptyFrame(rows, cols)

	frame.SetAlive(x, y)
	frame.SetAlive(x+1, y+1)
	frame.SetAlive(x+1, y+2)
	frame.SetAlive(x+2, y)
	frame.SetAlive(x+2, y+1)

	return frame
}

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

func (this *Frame) SetState(x, y int, state bool) {

	if y < len(this.Points) && y >= 0 {

		row := this.Points[y]

		if x < len(row) && x >= 0 {
			this.Points[y][x] = state
		}
	}
}

func (this *Frame) SetAlive(x, y int) {
	this.SetState(x, y, true)
}

func (this *Frame) SetDead(x, y int) {
	this.SetState(x, y, false)
}

func (this *Frame) GetAliveNeighbourCount(x, y int) int {

	aliveCount := 0

	for i := -1; i <= 1; i++ {

		if y+i < len(this.Points) && y+i >= 0 {

			for j := -1; j <= 1; j++ {

				if i != 0 || j != 0 {

					row := this.Points[y+i]

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

func (this *Frame) String() string {

	str := ""

	this.ForEachRow(func(row []bool, _ int) {

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

func (this *Frame) ForEachRow(callback func([]bool, int)) {

	for i, row := range this.Points {
		callback(row, i)
	}
}

func (this *Frame) ForEach(callback func(bool, int, int) bool) {

	this.ForEachRow(func(row []bool, y int) {

		for x, point := range row {

			continueLoop := callback(point, x, y)

			if !continueLoop {
				return
			}
		}
	})
}
