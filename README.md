# ebigrid

A flexible rectangular grid system for [Ebitengine](https://ebitengine.org/).

`ebigrid` provides a simple way to implement various types of rectangular grids.

## Installation

```bash
go get github.com/marble-fox/ebigrid
```

### Quick start

A quick overview of creating a grid. For more details, see the demos.

```go
// For finite grid, specify number of columns and rows;
// for infinite grid, set both values to 0
newGrid := grid.NewRectGrid(
  100,  // cell width
  100,  // cell height
  10,   // number of columns
  10,   // number of rows
)

// set spacing in pixels, if needed; both 0 by default
newGrid.HorizontalSpacing = 5
newGrid.VerticalSpacing = 5

// set incline as multiplier, if needed; both 0 by default
newGrid.InclineX = -1
newGrid.InclineY = 1
```

### Coordinate Conversion

```go
// Get cell coordinates from some position
cellX, cellY, exists := grid.GetCellCoordinates(posX, posY)

// Get position of origin of the cell
ox, py, exists := grid.GetCellOriginPosition(cellX, cellY)

// Get position in the center of the cell
cx, cy, exists := grid.GetCellCenterPosition(cellX, cellY)
```

## Demos

1. **Infinite Grid**
  <img width="1602" height="939" alt="image" src="https://github.com/user-attachments/assets/5fd689b2-a7dd-44f9-85ef-caa963fa2340" />

2. **Finite Grid**
  <img width="1602" height="939" alt="image" src="https://github.com/user-attachments/assets/64694e0a-afc1-4775-b696-af7d7fc08cb0" />

3. **Grid Spacing**
  <img width="1602" height="939" alt="image" src="https://github.com/user-attachments/assets/a490ec1b-d5f2-4e3a-9eea-868ceadfa546" />

4. **Diamond Grid**
  <img width="1602" height="939" alt="image" src="https://github.com/user-attachments/assets/a4624844-0fc5-44ca-a6cc-2103fdcc9ead" />

5. **Isometric Grid**
  <img width="1602" height="939" alt="image" src="https://github.com/user-attachments/assets/bde9f149-d607-4402-8003-48f2a28e0d48" />

6. **Any random arrangement of rectangles you want**
  <img width="1602" height="939" alt="image" src="https://github.com/user-attachments/assets/0f884d3a-9e6f-4fc4-89ec-1be9b9782560" />

Run them using:
```bash
go run ./demos/1_infiniteGrid
```
(Replace `1_infiniteGrid` with other demo folder names)

## TODO
- hexagonal grid
- triangular grid (maybe? does anyone even need this?)
