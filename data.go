package main

import "math"

// ---------- DATA TYPES ----------
type Point struct {
	x int
	y int
}

// Rooms: id -> (x,y)
var rooms = map[int]Point{
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
var links = [][2]int{
	{0, 4}, {0, 6}, {1, 3}, {4, 3},
	{5, 2}, {3, 5}, {4, 2}, {2, 1},
	{7, 6}, {7, 2}, {7, 4}, {6, 5},
}

// ---------- NORMALIZE ----------
var scale = 2

// do not start from big numbers !!!!!!!!!!!!!!! (do i even need them)
var minX, maxX = math.MaxInt, math.MinInt
var minY, maxY = math.MaxInt, math.MinInt
