package demos

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	grid "github.com/marble-fox/ebigrid"
)

type Game struct {
	Grid                     *grid.RectGrid
	CameraSpeed              int
	ResolutionX, ResolutionY int
	ShiftX, ShiftY           int
}

var DemoGame = Game{
	Grid:        nil,
	CameraSpeed: 3,
	ResolutionX: 1280,
	ResolutionY: 720,
	ShiftX:      0,
	ShiftY:      0,
}

func GameInitialize() {
	ebiten.SetWindowSize(DemoGame.ResolutionX, DemoGame.ResolutionY)
}

func (g *Game) Update() error {
	// Camera controls
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		g.ShiftX -= g.CameraSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		g.ShiftX += g.CameraSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.ShiftY -= g.CameraSpeed
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.ShiftY += g.CameraSpeed
	}

	// Incline controls
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		g.Grid.InclineY += 0.01
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		g.Grid.InclineY -= 0.01
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.Grid.InclineX += 0.01
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.Grid.InclineX -= 0.01
	}

	// Scale controls
	if ebiten.IsKeyPressed(ebiten.KeyMinus) && g.Grid.Scale > 0.25 {
		g.Grid.Scale -= 0.01
	}
	if ebiten.IsKeyPressed(ebiten.KeyEqual) && g.Grid.Scale < 10 {
		g.Grid.Scale += 0.01
	}

	// Cell size controls
	if ebiten.IsKeyPressed(ebiten.KeyI) {
		g.Grid.CellWidth += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyK) {
		g.Grid.CellWidth -= 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyO) {
		g.Grid.CellHeight += 1
	}
	if ebiten.IsKeyPressed(ebiten.KeyL) {
		g.Grid.CellHeight -= 1
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Grid.DebugDraw(screen, g.ShiftX, g.ShiftY, true)

	fps := ebiten.ActualFPS()
	tps := ebiten.ActualTPS()

	x, y := ebiten.CursorPosition()
	cellX, cellY, exists := g.Grid.CellCoordinatesFromPos(x+g.ShiftX, y+g.ShiftY)

	ebitenutil.DebugPrint(screen, fmt.Sprint(
		"FPS: ", fps,
		"\nTPS: ", tps,
		"\n",
		"\nCursor: ", x, y,
		"\nCell on cursor: ", cellX, cellY, exists,
		"\n",
		"\nCamera (WASD): ", g.ShiftX, g.ShiftY,
		"\nInclineX (Up/Down): ", g.Grid.InclineX,
		"\nInclineY (Left/Right): ", g.Grid.InclineY,
		"\nScale (-/+): ", g.Grid.Scale,
		"\nCellWidth (I/K): ", g.Grid.CellWidth,
		"\nCellHeight (O/L): ", g.Grid.CellHeight,
	))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}
