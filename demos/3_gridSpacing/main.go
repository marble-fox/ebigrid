package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/marble-fox/ebigrid"
	"github.com/marble-fox/ebigrid/demos"
)

func main() {
	newGrid := grid.NewRectGrid(
		50,
		50,
		10,
		5,
	)

	// Spacing
	newGrid.HorizontalSpacing = 5
	newGrid.VerticalSpacing = 15

	demos.DemoGame.Grid = newGrid

	demos.GameInitialize()
	err := ebiten.RunGame(&demos.DemoGame)
	if err != nil {
		log.Fatal(err)
	}
}
