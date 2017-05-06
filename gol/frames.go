package gol

type Frame struct {
	Points [][]int
}

func NewFrame(rows, cols int) *Frame {

	frame := &Frame{
		Points: make([][]int, rows, cols),
	}

	return frame
}

func (this *Frame) SetAlive(x, y int) {

}
