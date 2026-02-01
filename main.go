package main

import (
	"fmt"
)

// ---------- MAIN ----------
func main() {
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

	// ---------- DRAW ROOMS ----------
	for id, p := range pos {
		label := fmt.Sprintf("[%d]", id)
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
