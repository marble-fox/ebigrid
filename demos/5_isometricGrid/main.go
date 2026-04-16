package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/marble-fox/ebigrid"
	"github.com/marble-fox/ebigrid/demos"
)

func main() {
	// TODO comments for 45 angle isometric grid
	newGrid := grid.NewRectGrid(
		100,
		50,
		10,
		10,
	)

	newGrid.InclineX = -0.5
	newGrid.InclineY = 2.0

	demos.DemoGame.Grid = newGrid

	demos.GameInitialize()
	err := ebiten.RunGame(&demos.DemoGame)
	if err != nil {
		log.Fatal(err)
	}
}
