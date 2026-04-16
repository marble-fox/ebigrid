package grid

import (
	"fmt"
	"image"
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
	InclineY, InclineX                 float64
	Scale                              float64

	finite bool
}

func NewRectGrid(cellWidth, cellHeight, columns, rows int) *RectGrid {
	if columns < 0 || rows < 0 {
		panic(ErrInvalidGridSize)
	}

	zeroColumns := columns == 0
	zeroRows := rows == 0

	isFinite := !zeroColumns && !zeroRows

	// Is grid infinite?
	if !isFinite {
		if !(zeroColumns && zeroRows) {
			panic(ErrInvalidFiniteGridSize)
		}
	}

	return &RectGrid{
		Columns:           columns,
		Rows:              rows,
		CellWidth:         cellWidth,
		CellHeight:        cellHeight,
		HorizontalSpacing: 0,
		VerticalSpacing:   0,
		InclineY:          0,
		InclineX:          0,
		Scale:             1,

		finite: isFinite,
	}
}

func (g *RectGrid) CellCoordinatesFromPos(posX, posY int) (int, int, bool) {
	// I will not explain this, fuck you.

	if g.finite && (posX < 0 || posY < 0) {
		return 0, 0, false
	}

	pxf := float64(posX)
	pyf := float64(posY)

	fullWidth := float64(g.CellWidth+g.HorizontalSpacing) * g.Scale
	fullHeight := float64(g.CellHeight+g.VerticalSpacing) * g.Scale

	possibleX := math.Floor((pxf + g.InclineY*pyf) / fullWidth)
	possibleY := math.Floor((pyf + g.InclineX*pxf) / fullHeight)

	if g.finite && (possibleX >= float64(g.Columns) || possibleY >= float64(g.Rows)) {
		return 0, 0, false
	}

	dw := ((pxf + g.InclineY*pyf) / fullWidth) - possibleX
	dh := ((pyf + g.InclineX*pxf) / fullHeight) - possibleY

	rw := float64(g.CellWidth) / float64(g.CellWidth+g.HorizontalSpacing)
	rh := float64(g.CellHeight) / float64(g.CellHeight+g.VerticalSpacing)

	if dw > rw || dh > rh {
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

	fullWidth := float64(g.CellWidth+g.HorizontalSpacing) * g.Scale
	fullHeight := float64(g.CellHeight+g.VerticalSpacing) * g.Scale

	det := 1 - g.InclineY*g.InclineX
	if det == 0 {
		return 0, 0, false
	}

	x := (float64(coordsX)*fullWidth - g.InclineY*float64(coordsY)*fullHeight) / det
	y := (float64(coordsY)*fullHeight - g.InclineX*float64(coordsX)*fullWidth) / det

	return int(math.Floor(x)), int(math.Floor(y)), true
}

func (g *RectGrid) DebugDraw(screen *ebiten.Image, shiftX, shiftY int, drawCoordinates bool) {
	screenSize := screen.Bounds()
	sw, sh := screenSize.Dx(), screenSize.Dy()

	det := 1 - g.InclineY*g.InclineX
	if det == 0 {
		return
	}

	fullWidth := float64(g.CellWidth+g.HorizontalSpacing) * g.Scale
	fullHeight := float64(g.CellHeight+g.VerticalSpacing) * g.Scale

	var minCX, maxCX, minCY, maxCY int

	if g.finite {
		minCX, maxCX = 0, g.Columns-1
		minCY, maxCY = 0, g.Rows-1
	}

	corners := [][2]float64{
		{float64(shiftX), float64(shiftY)},
		{float64(shiftX + sw), float64(shiftY)},
		{float64(shiftX), float64(shiftY + sh)},
		{float64(shiftX + sw), float64(shiftY + sh)},
	}

	var fMinCX, fMaxCX, fMinCY, fMaxCY float64
	for i, c := range corners {
		cx := (c[0] + g.InclineY*c[1]) / fullWidth
		cy := (c[1] + g.InclineX*c[0]) / fullHeight
		if i == 0 {
			fMinCX, fMaxCX, fMinCY, fMaxCY = cx, cx, cy, cy
		} else {
			fMinCX = min(fMinCX, cx)
			fMaxCX = max(fMaxCX, cx)
			fMinCY = min(fMinCY, cy)
			fMaxCY = max(fMaxCY, cy)
		}
	}

	if g.finite {
		minCX = max(0, int(math.Floor(fMinCX)))
		maxCX = min(g.Columns-1, int(math.Floor(fMaxCX)))
		minCY = max(0, int(math.Floor(fMinCY)))
		maxCY = min(g.Rows-1, int(math.Floor(fMaxCY)))
	} else {
		minCX = int(math.Floor(fMinCX))
		maxCX = int(math.Floor(fMaxCX))
		minCY = int(math.Floor(fMinCY))
		maxCY = int(math.Floor(fMaxCY))
	}

	baseGeom := ebiten.GeoM{}
	baseGeom.SetElement(0, 0, (1/det)*g.Scale)
	baseGeom.SetElement(0, 1, (-g.InclineY/det)*g.Scale)
	baseGeom.SetElement(1, 0, (-g.InclineX/det)*g.Scale)
	baseGeom.SetElement(1, 1, (1/det)*g.Scale)

	op := &ebiten.DrawImageOptions{}

	debugImg := getCellShape(g.CellWidth, g.CellHeight)
	for cy := minCY; cy <= maxCY; cy++ {
		for cx := minCX; cx <= maxCX; cx++ {
			// Pre-calculate positions directly to avoid re-calculating fullWidth/fullHeight/det
			x := (float64(cx)*fullWidth - g.InclineY*float64(cy)*fullHeight) / det
			y := (float64(cy)*fullHeight - g.InclineX*float64(cx)*fullWidth) / det

			sx := x - float64(shiftX)
			sy := y - float64(shiftY)

			op.GeoM = baseGeom
			op.GeoM.Translate(sx, sy)

			screen.DrawImage(debugImg, op)

			if drawCoordinates {
				ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d\n%d", cx, cy), int(sx), int(sy))
			}
		}
	}
	debugImg.Deallocate()
}

func getCellShape(width, height int) *ebiten.Image {
	img := ebiten.NewImage(width, height)
	c := color.RGBA{R: 255, G: 0, B: 0, A: 255}

	// Draw borders using Fill
	// Top
	img.SubImage(image.Rect(0, 0, width, 1)).(*ebiten.Image).Fill(c)
	// Bottom
	img.SubImage(image.Rect(0, height-1, width, height)).(*ebiten.Image).Fill(c)
	// Left
	img.SubImage(image.Rect(0, 0, 1, height)).(*ebiten.Image).Fill(c)
	// Right
	img.SubImage(image.Rect(width-1, 0, width, height)).(*ebiten.Image).Fill(c)

	return img
}
