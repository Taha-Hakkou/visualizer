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

func drawMove(canvas [][]rune, x1, y1, x2, y2 int) {
	// 1) Horizontal segment
	if y1 == y2 {
		step := 1
		if x2 < x1 {
			step = -1
		}
		var tmp rune
		for x := x1 + 1; x != x2; x += step {
			if canvas[y1][x-step] == '•' { // or using: tmp != 0
				canvas[y1][x-step] = tmp
			}
			tmp = canvas[y1][x]
			canvas[y1][x] = '•' // cell shouldnt be empty because already drawn
		}
		return
	}

	// 2) Vertical segment
	if x1 == x2 {
		step := 1
		if y2 < y1 {
			step = -1
		}
		var tmp rune
		for y := y1 + 1; y != y2; y += step {
			if canvas[y-step][x2+1] == '•' { // or using: tmp != 0
				canvas[y-step][x2+1] = tmp
			}
			tmp = canvas[y][x2+1]
			canvas[y][x2+1] = '•' // cell shouldnt be empty because already drawn
		}
		return
	}

	// 3)
	// if y2 < y1 {
	// 	y1, y2 = y2, y1
	// 	x1, x2 = x2, x1
	// } // dont change direction !

	var x, y int = x1, y1

	var startWithPipes bool
	if y1+y2 > height || x1+x2 > width {
		// works in 3 quarters of the map
		// when using "and", will work only on 1 quarter
		startWithPipes = true
	}

	var tmp rune

	// 3.1
	if startWithPipes {
		// pipes
		abs := x2 - x1
		if abs < 0 {
			abs *= -1
		}
		for ; y2-y > abs; y++ { // is it a valid condition ?
			if canvas[y-1][x] == '•' {
				canvas[y-1][x] = tmp
			}
			tmp = canvas[y][x]
			canvas[y][x] = '•'
		}
	} else {
		// underscores
		if x2 > x1 {
			for ; x2-x != y2-y1; x++ {
				if canvas[y1][x] == '•' {
					canvas[y1][x] = tmp
				}
				tmp = canvas[y1][x+1]
				canvas[y1][x+1] = '•'
			}
		} else {
			for ; x2-x != y1-y2; x-- { // (x2-x)/(y2-y1) != -1
				if canvas[y1][x+1] == '•' {
					canvas[y1][x+1] = tmp
				}
				tmp = canvas[y1][x]
				canvas[y1][x] = '•'
			}
		}
	}

	// 3.2
	// slashes & backslashes
	step := 1
	if y1 > y2 {
		step = -1
	}
	if x2 > x1 {
		if tmp == '_' {
			x++
			y += step
			//
			// if canvas[y-step][x-1] == '•' {
			canvas[y-step][x-1] = tmp // sure it is a dot
			// }
			tmp = canvas[y][x]
			canvas[y][x] = '•'
		}
		for ; y != y2+1; x++ {
			canvas[y-step][x-1] = tmp // sure it is a dot
			tmp = canvas[y][x]
			canvas[y][x] = '•'
			// if canvas[y][x] == ' ' || canvas[y][x] == '_' {
			// 	canvas[y][x] = '\\'
			// }
			y += step
		}
	} else {
		// tx := x
		for ; y != y2+1; x-- {
			canvas[y-step][x+2] = tmp // sure it is a dot
			tmp = canvas[y][x+1]
			canvas[y][x+1] = '•'
			// if canvas[y][x+1] == ' ' || canvas[y][x+1] == '_' && tx != x {
			// 	canvas[y][x+1] = '/'
			// }
			y += step
		}
	}

	// 3.3
	if startWithPipes {
		// underscores
		if x2 > x1 {
			// x-- // no need because just a move !
			// tmp = canvas[y][x]
			for ; x < x2; x++ {
				canvas[y][x] = tmp // where the dot is
				tmp = canvas[y][x+1]
				canvas[y][x+1] = '•'
				// flush

				// canvas[y][x] = tmp // sure it is a dot
				// if canvas[y2][x] == ' ' {
				// 	canvas[y2][x] = '_'
				// }
			}
		} else {
			// x++
			for ; x > x2; x-- {
				canvas[y][x] = tmp // where the dot is
				tmp = canvas[y][x-1]
				canvas[y][x-1] = '•'
				// flush

				// tmp = canvas[y][x+1]
				// canvas[y][x+1] = '•'
				// if canvas[y2][x] == ' ' {
				// 	canvas[y2][x] = '_'
				// }
			}
		}
	} else {
		// pipes
		for ; y < y2; y++ {
			canvas[y][x] = tmp // where the dot is
			tmp = canvas[y+1][x]
			canvas[y+1][x] = '•'
			// flush

			// if canvas[y][x] == ' ' {
			// 	canvas[y][x] = '|'
			// }
		}
	}
}
