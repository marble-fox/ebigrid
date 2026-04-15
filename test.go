package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/marble-fox/ebigrid/grid"
)

type Game struct {
	Grid                     *grid.RectGrid
	cameraSpeed              int
	resolutionX, resolutionY int
	shiftX, shiftY           int
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.shiftX -= g.cameraSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.shiftX += g.cameraSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.shiftY -= g.cameraSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.shiftY += g.cameraSpeed
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Grid.DebugDraw(screen, g.shiftX, g.shiftY, true)

	fps := ebiten.ActualFPS()
	tps := ebiten.ActualTPS()

	x, y := ebiten.CursorPosition()
	cellX, cellY, exists := g.Grid.CellCoordinatesFromPos(x+g.shiftX, y+g.shiftY)

	ebitenutil.DebugPrint(screen, fmt.Sprint(
		"FPS: ", fps,
		"\nTPS: ", tps,
		"\nCursor: ", x, y,
		"\nCell: ", cellX, cellY, exists,
	))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func main() {
	newGrid, err := grid.NewRectGrid(
		42,
		42,
		0,
		0,
		15,
		15,
	)
	if err != nil {
		log.Fatal(err)
	}

	game := Game{
		Grid:        newGrid,
		cameraSpeed: 3,
		resolutionX: 1280,
		resolutionY: 720,
		shiftX:      0,
		shiftY:      0,
	}

	ebiten.SetWindowSize(game.resolutionX, game.resolutionY)
	//ebiten.SetFullscreen(true)

	err = ebiten.RunGame(&game)
	if err != nil {
		log.Fatal(err)
	}
}
