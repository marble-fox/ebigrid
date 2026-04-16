package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/marble-fox/ebigrid"
	"github.com/marble-fox/ebigrid/demos"
)

func main() {
	newGrid := grid.NewRectGrid(
		100,
		100,
		10,
		10,
	)

	// To create a diamond grid, we skew both axes.
	// For a 45-degree diamond effect, we use a skew factor of 1 (since tan(45°) = 1).
	// This means for every 1 unit of movement on one axis, we shift 1 unit on the other.

	// Skew the X-axis (vertical lines)
	newGrid.InclineX = -1
	// Skew the Y-axis (horizontal lines) in the opposite direction
	newGrid.InclineY = 1

	// If the explanation isn't clear enough, try playing around with the incline values in real time
	// by using the arrow keys (Up/Down for InclineX, Left/Right for InclineY).
	// In particular, try holding Down+Left or Up+Right to "rotate" the grid clockwise or counterclockwise.

	demos.DemoGame.Grid = newGrid

	demos.GameInitialize()
	err := ebiten.RunGame(&demos.DemoGame)
	if err != nil {
		log.Fatal(err)
	}
}
