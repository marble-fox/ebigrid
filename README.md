# ebigrid

A flexible rectangular grid system for [Ebitengine](https://ebitengine.org/).

`ebigrid` provides a simple way to implement various types of rectangular grids.

## Installation

```bash
go get github.com/marble-fox/ebigrid
```

### Coordinate Conversion

```go
// Get cell coordinates from some position
cellX, cellY, exists := grid.CellCoordinatesFromPos(posX, posY)

// Get screen position from cell coordinates
x, y, exists := grid.CellPosFromCoordinates(cellX, cellY)
```

## Demos

The project includes several demos showcasing different grid configurations:

1. **Infinite Grid**
<img width="1602" height="939" alt="image" src="https://github.com/user-attachments/assets/70f5c52d-3c11-4333-a12d-e34b67305cef" />

2. **Finite Grid**
<img width="1602" height="939" alt="image" src="https://github.com/user-attachments/assets/c544237e-9131-408f-9b0b-dc0fcd92200a" />

3. **Grid Spacing**
<img width="1602" height="939" alt="image" src="https://github.com/user-attachments/assets/d4916299-9c25-4014-9e7a-c0e9cdc22d02" />

4. **Diamond Grid**
<img width="1602" height="939" alt="image" src="https://github.com/user-attachments/assets/3b3b00a1-b5c0-4054-abbb-f2b0ab05b4d8" />

5. **Isometric Grid**
<img width="1602" height="939" alt="image" src="https://github.com/user-attachments/assets/d8d8bf56-ceb8-4571-8428-de0793a1137c" />

6. **Any random arrangement of rectangles you want**
<img width="1602" height="939" alt="image" src="https://github.com/user-attachments/assets/0a71913d-22c1-4243-b4ed-cff86dfd5e6d" />


Run them using:
```bash
go run ./demos/1_infiniteGrid
```
(Replace `1_infiniteGrid` with other demo folder names)

## TODO
- a function that returns the position at the center of a cell
- hexagonal grid
- triangular grid (maybe? does anyone even need this?)
