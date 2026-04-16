package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/marble-fox/ebigrid"
	"github.com/marble-fox/ebigrid/demos"
)

func main() {
	// To create a finite grid, specify greater than 0 values for columns and rows
	newGrid := grid.NewRectGrid(
		50,
		50,
		10,
		10,
	)

	demos.DemoGame.Grid = newGrid

	demos.GameInitialize()
	err := ebiten.RunGame(&demos.DemoGame)
	if err != nil {
		log.Fatal(err)
	}
}
