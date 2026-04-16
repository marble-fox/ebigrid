package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/marble-fox/ebigrid"
	"github.com/marble-fox/ebigrid/demos"
)

func main() {
	// To create an infinite grid, the number of columns and rows must be set to 0
	newGrid := grid.NewRectGrid(
		50,
		50,
		0,
		0,
	)

	demos.DemoGame.Grid = newGrid

	demos.GameInitialize()
	err := ebiten.RunGame(&demos.DemoGame)
	if err != nil {
		log.Fatal(err)
	}
}
