package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	paths := getPaths()
	x, y, _ := getBounds(paths)
	c := getScan(paths, x, y)

	done := false
	sand := 0
	for !done {
		sand++
		done = c.addSand()
	}

	fmt.Printf("day14 part1: %d\n", sand-1)
}

func getPaths() [][][]int {
	file, _ := os.Open("input.txt")
	fs := bufio.NewScanner(file)
	fs.Split(bufio.ScanLines)

	paths := make([][][]int, 0)

	for fs.Scan() {
		stub := make([][]int, 0)
		path := strings.Split(fs.Text(), " -> ")

		for _, p := range path {
			points := strings.Split(p, ",")
			x, _ := strconv.Atoi(points[0])
			y, _ := strconv.Atoi(points[1])
			stub = append(stub, []int{x, y})
		}
		paths = append(paths, stub)
	}
	return paths
}

func getBounds(paths [][][]int) (int, int, int) {
	maxX := 0
	maxY := 0
	minX := 1000

	for i := range paths {
		for j := range paths[i] {
			if paths[i][j][0] > maxX {
				maxX = paths[i][j][0]
			}
			if paths[i][j][0] < minX {
				minX = paths[i][j][0]
			}
			if paths[i][j][1] > maxY {
				maxY = paths[i][j][1]
			}
		}
	}
	return maxX, maxY, minX
}

func getScan(paths [][][]int, maxX, maxY int) *cave {
	c := new(cave)
	c.maxY = maxY

	c.scan = make([][]string, maxY+1)
	for i := range c.scan {
		c.scan[i] = make([]string, maxX+1)
	}

	for _, path := range paths {
		for i := 0; i < len(path)-1; i++ {
			startY := path[i][1]
			endY := path[i+1][1]

			startX := path[i][0]
			endX := path[i+1][0]

			if startY > endY {
				startY, endY = endY, startY
			}

			if startX > endX {
				startX, endX = endX, startX
			}

			for y := startY; y <= endY; y++ {
				for x := startX; x <= endX; x++ {
					c.scan[y][x] = "#"
				}
			}
		}
	}

	return c
}

type cave struct {
	maxY int
	scan [][]string
}

func (c *cave) addSand() bool {
	y, x := evalDrop(c.scan, 0, 500)
	if y == c.maxY {
		return true
	}
	c.scan[y][x] = "#"
	return false
}

func evalDrop(cave [][]string, y, x int) (int, int) {
	switch {
	case y == len(cave)-1:
		return y, x
	case cave[y+1][x] != "#":
		return evalDrop(cave, y+1, x)
	case x > 1 && cave[y+1][x-1] != "#":
		return evalDrop(cave, y+1, x-1)
	case x < len(cave[0])-2 && cave[y+1][x+1] != "#":
		return evalDrop(cave, y+1, x+1)
	}
	return y, x
}
