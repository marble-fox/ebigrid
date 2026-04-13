package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/marble-fox/ebigrid/grid"
)

type Game struct {
	Grid           *grid.RectGrid
	shiftX, shiftY int
}

func (g *Game) Update() error {
	speed := 3
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.shiftX -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.shiftX += speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.shiftY -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.shiftY += speed
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
	game := Game{
		Grid: grid.NewRectGrid(
			30,
			30,
			0,
			0,
			30,
			30,
		),
		shiftX: 0,
		shiftY: 0,
	}

	ebiten.SetWindowSize(1296, 810)
	//ebiten.SetFullscreen(true)

	err := ebiten.RunGame(&game)
	if err != nil {
		log.Fatal(err)
	}
}
