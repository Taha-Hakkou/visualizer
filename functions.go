package main

func drawLine(canvas [][]rune, x1, y1, x2, y2 int) {

	// 1) Horizontal segment
	if y1 == y2 {
		step := 1
		if x2 < x1 {
			step = -1
		}
		for x := x1; x != x2; x += step {
			if canvas[y1][x] == ' ' {
				canvas[y1][x] = '-'
			}
		}
		return
	}

	// 2) Vertical segment
	if x1 == x2 {
		step := 1
		if y2 < y1 {
			step = -1
		}
		for y := y1 + step; y != y2+step; y += step {
			if canvas[y][x2+1] == ' ' {
				canvas[y][x2+1] = '|'
			}
		}
		return
	}

	// 3)
	// underscores (in case starting with underscores - can be pipes)
	xstep := 1
	if x2 < x1 {
		xstep = -1
	}
	x := x1
	for ; (x2-x)/(y2-y1) != 0; x += xstep { // x != x2
		if canvas[y1][x] == ' ' {
			canvas[y1][x] = '_'
		}
	}

	// slashes & backslashes
	// y := y1
	// ystep := xstep
	// if x2-x == y-y2 {
	// 	ystep = -xstep
	// }
	// for ; x != x2; x += xstep { // (x2-x)/(y2-y) != 0
	// 	if canvas[y][x] == ' ' {
	// 		canvas[y][x] = '/'
	// 	}
	// 	y += ystep // can be different of x (and use backslash)
	// }

	// pipes

	// // 3) Vertical or diagonal turn
	// if y1 == y2 {
	// 	return
	// }

	// // draw the turn
	// turnChar := '/'
	// if (x2-x1)*(y2-y1) > 0 {
	// 	turnChar = '\\'
	// }

	// if canvas[y1][x2] == ' ' {
	// 	canvas[y1][x2] = turnChar
	// }

}
