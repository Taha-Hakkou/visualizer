package main

func hyphens(canvas [][]rune, y, x1, x2 int) {
	// if x2 < x1 {
	// 	x1, x2 = x2, x1
	// }
	var step int = 1
	if x2 < x1 {
		step = -1
	}
	for x := x1 + step; (x2-x)*step > 0; x += step {
		// check because room name takes more than a cell
		if canvas[y][x] == ' ' {
			canvas[y][x] = '-'
		}
	}
}

func pipes(canvas [][]rune, x, y1, y2 int) {
	// if y2 < y1 {
	// 	y1, y2 = y2, y1
	// }
	var step int = 1
	if y2 < y1 {
		step = -1
	}
	for y := y1 + step; (y2-y)*step > 0; y += step {
		// moved 1 cell to the right to be appropriately aligned
		// fmt.Println(x, y, y1)
		if canvas[y][x] == ' ' {
			canvas[y][x] = '|'
		}
	}
}
