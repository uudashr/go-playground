package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math/rand/v2"
	"os"
)

type Dot struct {
	X, Y    int
	IsPlace bool   // True if this dot is a place
	PlaceID string // ID of the place if it is one
}

type Map struct {
	Grid   [][]*Dot
	Width  int
	Height int
	Places []*Dot
}

func NewMap(width, height int) *Map {
	grid := make([][]*Dot, height)
	for i := range grid {
		grid[i] = make([]*Dot, width)
	}
	return &Map{Grid: grid, Width: width, Height: height, Places: []*Dot{}}
}

func (m *Map) AddPlace(x, y int, id string) {
	dot := &Dot{X: x, Y: y, IsPlace: true, PlaceID: id}
	m.Grid[y][x] = dot
	m.Places = append(m.Places, dot)
}

func (m *Map) AddPath(start, end *Dot) {
	currentX, currentY := start.X, start.Y
	endX, endY := end.X, end.Y

	// Use a simple pathfinding approach to move towards the destination
	for currentX != endX || currentY != endY {
		m.Grid[currentY][currentX] = &Dot{X: currentX, Y: currentY, IsPlace: false}

		if currentX < endX {
			currentX++
		} else if currentX > endX {
			currentX--
		}

		if currentY < endY {
			currentY++
		} else if currentY > endY {
			currentY--
		}
	}

	// Mark the end of the route as well
	m.Grid[endY][endX] = end
}

func (m *Map) AddRoute(start, end *Dot) {
	dx := end.X - start.X
	dy := end.Y - start.Y

	steps := max(abs(dx), abs(dy))
	for i := 1; i <= steps; i++ {
		x := start.X + i*dx/steps
		y := start.Y + i*dy/steps
		if m.Grid[y][x] == nil {
			m.Grid[y][x] = &Dot{X: x, Y: y, IsPlace: false}
		}
	}
}

func (m *Map) DrawMap(filename string) error {
	imgWidth := m.Width * 20
	imgHeight := m.Height * 20

	img := image.NewRGBA(image.Rect(0, 0, imgWidth, imgHeight))

	// Background color
	bgColor := color.RGBA{248, 247, 247, 255}

	draw.Draw(img, img.Bounds(), &image.Uniform{bgColor}, image.Point{}, draw.Src)

	// Colors
	roadColor := color.RGBA{173, 185, 199, 255} // roads
	placeColor := color.RGBA{255, 0, 0, 255}    // places

	// Draw roads
	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			if m.Grid[y][x] != nil && !m.Grid[y][x].IsPlace {
				rect := image.Rect(x*20, y*20, (x+1)*20, (y+1)*20)
				draw.Draw(img, rect, &image.Uniform{roadColor}, image.Point{}, draw.Src)
			}
		}
	}

	// Draw places as larger dots or circles
	for _, place := range m.Places {
		drawCircle(img, place.X*20+10, place.Y*20+10, 30, placeColor)
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	return png.Encode(file, img)
}

func drawCircle(img *image.RGBA, x, y, radius int, col color.Color) {
	for dx := -radius; dx <= radius; dx++ {
		for dy := -radius; dy <= radius; dy++ {
			if dx*dx+dy*dy <= radius*radius {
				img.Set(x+dx, y+dy, col)
			}
		}
	}
}

func (m *Map) GenerateRandomPlaces(numPlaces int) {
	for i := 0; i < numPlaces; i++ {
		x := rand.IntN(m.Width)
		y := rand.IntN(m.Height)
		placeID := fmt.Sprintf("P%d", i+1)
		m.AddPlace(x, y, placeID)
	}
}

func (m *Map) GenerateRoutes() {
	for i := 0; i < len(m.Places)-1; i++ {
		m.AddPath(m.Places[i], m.Places[i+1])
	}

	// Optionally connect the last place to the first to make a loop
	m.AddPath(m.Places[len(m.Places)-1], m.Places[0])
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// User-defined map size and number of places
	mapWidth := 500
	mapHeight := 500
	numPlaces := 30

	m := NewMap(mapWidth, mapHeight)

	// Generate random places and routes
	m.GenerateRandomPlaces(numPlaces)
	m.GenerateRoutes()

	// Draw and save the map as an image
	if err := m.DrawMap("generated_map.png"); err != nil {
		fmt.Println("Error drawing map:", err)
	} else {
		fmt.Println("Map saved as generated_map.png")
	}
}
