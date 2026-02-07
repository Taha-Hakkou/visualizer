package main

func hyphens(canvas [][]rune, y, x1, x2 int) {
	if x2 < x1 {
		x1, x2 = x2, x1
	}
	for x := x1 + 1; x < x2; x++ {
		// check because room name takes more than a cell
		if canvas[y][x] == ' ' {
			canvas[y][x] = '-'
		}
	}
}

func pipes(canvas [][]rune, x, y1, y2 int) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	for y := y1 + 1; y < y2; y++ {
		// moved 1 cell to the right to be appropriately aligned
		if canvas[y][x+1] == ' ' {
			canvas[y][x+1] = '|'
		}
	}
}
