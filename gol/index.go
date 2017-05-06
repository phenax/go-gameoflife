package gol

type GameOfLife struct {
	Rows    int
	Columns int
}

func NewGameOfLife(rows, cols int) *GameOfLife {

	game := &GameOfLife{
		Rows:    rows,
		Columns: cols,
	}

	return game
}

func (this *GameOfLife) LoadFrame(frame *Frame) {

}

func (this *GameOfLife) StartLoop() {

}
