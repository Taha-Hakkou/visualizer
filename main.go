package main

import (
	"fmt"
)

var width, height int = 0, 0

// ---------- MAIN ----------
func main() {
	getData()

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

	pos := make(map[string]Room)
	for id, p := range rooms {
		gx := (p.x - minX) * scale
		gy := (p.y - minY) * scale // to invert Y, use (maxY - p.y) instead
		pos[id] = Room{gx, gy}     // is id converted implicitly ???
	}

	// ---------- CANVAS ----------
	for _, p := range pos {
		if p.x > width {
			width = p.x
		}
		if p.y > height {
			height = p.y
		}
	}
	// min to add. why ????
	width += 3
	height += 1 // adding to width and height in case some lines go beyond !

	canvas := make([][]rune, height)
	for i := range canvas {
		canvas[i] = make([]rune, width)
		for j := range canvas[i] {
			canvas[i][j] = ' '
		}
	}
	// can i populate them in single line
	// + add space between lines so that slaches andd backslaches align properly !

	// ---------- DRAW ROOMS ----------
	for id, p := range pos {
		label := fmt.Sprintf("[%s]", id)
		for i, ch := range label {
			canvas[p.y][p.x+i] = ch
		}
	}

	// ---------- DRAW LINES ----------
	for _, e := range links {
		a, b := e[0], e[1]
		p1, p2 := pos[a], pos[b]
		drawLine(canvas, p1.x, p1.y, p2.x, p2.y)
	}

	// ---------- PRINT ----------
	for _, row := range canvas {
		fmt.Println(string(row))
	}
}
