package main

func drawLine(canvas [][]rune, x1, y1, x2, y2 int) {
	// 1) Horizontal segment
	if y1 == y2 {
		step := 1
		if x2 < x1 {
			step = -1
		}
		for x := x1 + 1; x != x2; x += step {
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
		for y := y1 + 1; y != y2; y += step {
			if canvas[y][x2+1] == ' ' {
				canvas[y][x2+1] = '|'
			}
		}
		return
	}

	// 3)
	if y2 < y1 {
		y1, y2 = y2, y1
		x1, x2 = x2, x1
	} // now y1 is always bigger than y2

	var x, y int = x1, y1

	var startWithPipes bool
	if y1+y2 > height || x1+x2 > width {
		// works in 3 quarters of the map
		// when using "and", will work only on 1 quarter
		startWithPipes = true
	}

	// 3.1
	if startWithPipes {
		// pipes
		abs := x2 - x1
		if abs < 0 {
			abs *= -1
		}
		for ; y2-y > abs; y++ { // is it a valid condition ?
			if canvas[y][x] == ' ' {
				canvas[y][x] = '|'
			}
		}
	} else {
		// underscores
		if x2 > x1 {
			for ; x2-x != y2-y1; x++ {
				if canvas[y1][x+1] == ' ' {
					canvas[y1][x+1] = '_'
				}
			}
		} else {
			for ; x2-x != y1-y2; x-- { // (x2-x)/(y2-y1) != -1
				if canvas[y1][x] == ' ' {
					canvas[y1][x] = '_'
				}
			}
		}
	}

	// 3.2
	// slashes & backslashes
	if x2 > x1 {
		if canvas[y][x] == '_' {
			x++
			if y2 > y1 {
				y++
			} else {
				y--
			}
		}
		for ; y != y2+1; x++ {
			if canvas[y][x] == ' ' || canvas[y][x] == '_' {
				canvas[y][x] = '\\'
			}
			if y2 > y1 {
				y++
			} else {
				y--
			}
		}
	} else {
		tx := x
		for ; y != y2+1; x-- {
			if canvas[y][x+1] == ' ' || canvas[y][x+1] == '_' && tx != x {
				canvas[y][x+1] = '/'
			}
			if y2 > y1 {
				y++
			} else {
				y--
			}
		}
	}

	// 3.3
	if startWithPipes {
		// underscores
		if x2 > x1 {
			x--
			for ; x < x2; x++ {
				if canvas[y2][x] == ' ' {
					canvas[y2][x] = '_'
				}
			}
		} else {
			x++
			for ; x > x2; x-- {
				if canvas[y2][x] == ' ' {
					canvas[y2][x] = '_'
				}
			}
		}
	} else {
		// pipes
		for ; y < y2; y++ {
			if canvas[y][x] == ' ' {
				canvas[y][x] = '|'
			}
		}
	}

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
