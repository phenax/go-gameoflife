
# Go-GameOfLife
Conway's game of life implementation in golang


### Instructions
- Run `go get github.com/phenax/go-gameoflife` to install
- Run `go-gameoflife` to see stuff move


### Available lives
- gol.NewEmptyFrame      - Creates an empty frame. You can use frame.SetAlive(x, y int) to change cells
- gol.NewBlinkerFrame    - `|` becomes `---` and back to `|` forever
- gol.NewGliderFrame     - Create a glider life(google it)
- gol.NewSpaceshipFrame  - Create a Lightweight spaceship life(google it)(which doesnt look anything like a spaceship btw)
- gol.NewPulsarFrame     - Create a pulsar life(google it)(My personal favorite)
