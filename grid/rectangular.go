package grid

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var ErrInvalidGridSize = fmt.Errorf("invalid grid size")
var ErrInvalidFiniteGridSize = fmt.Errorf("invalid finite grid size")

type RectGrid struct {
	CellWidth, CellHeight              int
	HorizontalSpacing, VerticalSpacing int
	Columns, Rows                      int

	rw, rh         float64
	finite         bool
	cellDebugShape *ebiten.Image
}

func NewRectGrid(cellWidth, cellHeight, horizontalSpacing, verticalSpacing, columns, rows int) (*RectGrid, error) {
	if columns < 0 || rows < 0 {
		return nil, ErrInvalidGridSize
	}

	rw := float64(cellWidth) / float64(cellWidth+horizontalSpacing)
	rh := float64(cellHeight) / float64(cellHeight+verticalSpacing)

	debugImg := getCellShape(cellWidth, cellHeight)

	zeroColumns := columns == 0
	zeroRows := rows == 0

	// Is grid infinite?
	if zeroColumns && zeroRows {
		return &RectGrid{
			Columns:           columns,
			Rows:              rows,
			CellWidth:         cellWidth,
			CellHeight:        cellHeight,
			HorizontalSpacing: horizontalSpacing,
			VerticalSpacing:   verticalSpacing,

			rw:             rw,
			rh:             rh,
			finite:         false,
			cellDebugShape: debugImg,
		}, nil
	}

	if zeroColumns || zeroRows {
		return nil, ErrInvalidFiniteGridSize
	}

	// Finite grid
	return &RectGrid{
		Columns:           columns,
		Rows:              rows,
		CellWidth:         cellWidth,
		CellHeight:        cellHeight,
		HorizontalSpacing: horizontalSpacing,
		VerticalSpacing:   verticalSpacing,

		rw:             rw,
		rh:             rh,
		finite:         true,
		cellDebugShape: debugImg,
	}, nil
}

func (g *RectGrid) CellCoordinatesFromPos(posX, posY int) (int, int, bool) {
	// I will not explain this, fuck you.

	if g.finite && (posX < 0 || posY < 0) {
		return 0, 0, false
	}

	pxf := float64(posX)
	pyf := float64(posY)

	fullWidth := float64(g.CellWidth + g.HorizontalSpacing)
	fullHeight := float64(g.CellHeight + g.VerticalSpacing)

	possibleX := math.Floor(pxf / fullWidth)
	possibleY := math.Floor(pyf / fullHeight)

	if g.finite && (possibleX >= float64(g.Columns) || possibleY >= float64(g.Rows)) {
		return 0, 0, false
	}

	dw := (pxf / fullWidth) - possibleX
	dh := (pyf / fullHeight) - possibleY

	if dw > g.rw || dh > g.rh {
		return 0, 0, false
	}

	return int(possibleX), int(possibleY), true
}

func (g *RectGrid) CellPosFromCoordinates(coordsX, coordsY int) (int, int, bool) {
	if g.finite {
		if coordsX < 0 || coordsY < 0 || coordsX >= g.Columns || coordsY >= g.Rows {
			return 0, 0, false
		}
	}

	x := coordsX*g.HorizontalSpacing + coordsX*g.CellWidth
	y := coordsY*g.VerticalSpacing + coordsY*g.CellHeight

	return x, y, true
}

func (g *RectGrid) DebugDraw(screen *ebiten.Image, shiftX, shiftY int, drawCoordinates bool) {
	screenSize := screen.Bounds()

	fullWidth := g.CellWidth + g.HorizontalSpacing
	fullHeight := g.CellHeight + g.VerticalSpacing

	// Determine range of visible cells
	minCellX := int(math.Floor(float64(shiftX) / float64(fullWidth)))
	minCellY := int(math.Floor(float64(shiftY) / float64(fullHeight)))

	maxCellX := int(math.Floor(float64(shiftX+screenSize.Dx()) / float64(fullWidth)))
	maxCellY := int(math.Floor(float64(shiftY+screenSize.Dy()) / float64(fullHeight)))

	for cy := minCellY; cy <= maxCellY; cy++ {
		for cx := minCellX; cx <= maxCellX; cx++ {
			// World position of cell
			wx, wy, exists := g.CellPosFromCoordinates(cx, cy)
			if !exists {
				continue
			}

			// Screen position (world - camera)
			sx := wx - shiftX
			sy := wy - shiftY

			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(sx), float64(sy))

			if drawCoordinates {
				// To draw the text label on top of the cell, we must create a copy of the debug shape image.
				// ebitenutil.DebugPrint draws text directly into an image, so we use a copy for each cell.
				img := ebiten.NewImageFromImage(g.cellDebugShape)
				ebitenutil.DebugPrint(img, fmt.Sprint(cx, "\n", cy))
				screen.DrawImage(img, op)
				img.Clear()
				continue
			}

			screen.DrawImage(g.cellDebugShape, op)
		}
	}
}

func getCellShape(width, height int) *ebiten.Image {
	img := ebiten.NewImage(width, height)

	c := color.RGBA{R: 255, G: 0, B: 0, A: 255}
	for x := 0; x < width; x++ {
		img.Set(x, 0, c)
		img.Set(x, height-1, c)
	}

	for y := 0; y < height; y++ {
		img.Set(0, y, c)
		img.Set(width-1, y, c)
	}

	return img
}
