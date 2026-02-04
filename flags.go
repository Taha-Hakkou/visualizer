package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func getData() {
	// check whether stdin is coming from a pipe/file or from a terminal (TTY)
	stat, _ := os.Stdin.Stat()
	isPiped := (stat.Mode() & os.ModeCharDevice) == 0

	if !isPiped || len(os.Args) > 1 {
		fmt.Println("Usage: ./lem-in [FILE] | ./visualizer")
		return
	}

	bytes, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err) // why panic ?
	}
	lines := strings.Split(string(bytes), "\n")

	// need to check global error of lem-in

	// Number of ants
	// n, _ = strconv.Atoi(lines[0])

	var i int
	// var start, end string

	// recheck regex: surrounding spaces
	roomex := regexp.MustCompile(`^([A-Za-z0-9]+) ([0-9]+) ([0-9]+)$`)
	for i = 1; i < len(lines); i++ {
		if roomex.MatchString(lines[i]) { // contains and not match: resolved
			groups := roomex.FindStringSubmatch(lines[i])
			x, _ := strconv.Atoi(groups[2])
			y, _ := strconv.Atoi(groups[3])
			rooms[groups[1]] = Room{x, y}
			// fmt.Println(lines[i])
		} else if lines[i] == "##start" {
			// start = lines[i+1]
		} else if lines[i] == "##end" {
			// end = lines[i+1]
		} else {
			break
		}
	}

	// 	fmt.Println()

	tunlex := regexp.MustCompile(`^([A-Za-z0-9]+)-([A-Za-z0-9]+)$`)
	for ; i < len(lines); i++ {
		if tunlex.MatchString(lines[i]) { // contains and not match: resolved
			groups := tunlex.FindStringSubmatch(lines[i])
			links = append(links, [2]string{groups[1], groups[2]})
			// fmt.Println(lines[i])
		} else {
			break
		}
	}

	// 	// movex := regexp.MustCompile(`^L[0-9]-[0-9]( (L[0-9]-[0-9])*)$`)
	// 	// for ; i < len(lines); i++ {
	// 	// 	if movex.MatchString(lines[i]) { // contains and not match
	// 	// 		fmt.Println(lines[i])
	// 	// 	}
	// 	// }

	// fmt.Println(n)
	// fmt.Println(start)
	// fmt.Println(end)
}
