package main

import (
	"fmt"
	"math"
)

// ---------- DATA TYPES ----------
type Point struct {
	x int
	y int
}

// ---------- MAIN ----------
func main() {

	// Rooms: id -> (x,y)
	rooms := map[int]Point{
		0: {9, 5},
		1: {23, 3},
		2: {16, 7},
		3: {16, 3},
		4: {16, 5},
		5: {9, 3},
		6: {1, 5},
		7: {4, 8},
	}

	// Links
	links := [][2]int{
		{0, 4}, {0, 6}, {1, 3}, {4, 3},
		{5, 2}, {3, 5}, {4, 2}, {2, 1},
		{7, 6}, {7, 2}, {7, 4}, {6, 5},
	}

	// ---------- NORMALIZE ----------
	scale := 3


	// do not start from big numbers !!!!!!!!!!!!!!! (do i even need them)
	minX, maxX := math.MaxInt, math.MinInt
	minY, maxY := math.MaxInt, math.MinInt

	for _, p := range rooms {
		if p.x < minX {
			minX = p.x
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.y > maxY {
			maxY = p.y
		}
	}

	pos := make(map[int]Point)
	for id, p := range rooms {
		gx := (p.x - minX) * scale
		gy := (p.y - minY) * scale // to invert Y, use (maxY - p.y) instead
		pos[id] = Point{gx, gy}
	}

	// ---------- CANVAS ----------
	width, height := 0, 0
	for _, p := range pos {
		if p.x > width {
			width = p.x
		}
		if p.y > height {
			height = p.y
		}
	}
	width += 10
	height += 5 // adding to width and height in case some lines go beyond ! 

	canvas := make([][]rune, height)
	for i := range canvas {
		canvas[i] = make([]rune, width)
		for j := range canvas[i] {
			canvas[i][j] = ' '
		}
	}
	// can i populate them in single line
	// + add space between lines so that slaches andd backslaches align properly !

	// ---------- DRAW LINES ----------
	drawLine := func(x1, y1, x2, y2 int) {
		dx := x2 - x1
		dy := y2 - y1
		steps := int(math.Max(math.Abs(float64(dx)), math.Abs(float64(dy))))

		for i := 0; i <= steps; i++ {
			x := int(math.Round(float64(x1) + float64(dx)*float64(i)/float64(steps)))
			y := int(math.Round(float64(y1) + float64(dy)*float64(i)/float64(steps)))

			var ch rune
			switch {
			case dx == 0:
				ch = '|'
			case dy == 0:
				ch = '-'
			case (dx > 0) == (dy > 0):
				ch = '\\'
			default:
				ch = '/'
			}

			if canvas[y][x] == ' ' {
				canvas[y][x] = ch
			}
		}
	}

	for _, e := range links {
		a, b := e[0], e[1]
		p1, p2 := pos[a], pos[b]
		drawLine(p1.x, p1.y, p2.x, p2.y)
	}

	// ---------- DRAW ROOMS ----------
	for id, p := range pos {
		label := fmt.Sprintf("[%d]", id)
		for i, ch := range label {
			canvas[p.y][p.x+i] = ch
		}
	}

	// ---------- PRINT ----------
	for _, row := range canvas {
		fmt.Println(string(row))
	}
}

