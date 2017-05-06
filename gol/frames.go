package gol

type Frame struct {
	Points [][]bool
}

func EmptyFrame(rows, cols int) *Frame {

	frame := &Frame{
		Points: make([][]bool, rows, rows),
	}

	for i := range frame.Points {
		frame.Points[i] = make([]bool, cols, cols)
	}

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
