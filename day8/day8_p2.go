package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	grid := make([][]int, 0)

	for fileScanner.Scan() {
		row := make([]int, 0)
		trees := strings.Split(fileScanner.Text(), "")

		for _, t := range trees {
			height, _ := strconv.Atoi(t)
			row = append(row, height)
		}
		grid = append(grid, row)
	}

	var bestScore int

	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[y])-1; x++ {
			up := viewUp(grid, y, x)
			down := viewDown(grid, y, x)
			left := viewLeft(grid, y, x)
			right := viewRight(grid, y, x)

			score := up * down * left * right
			if score > bestScore {
				bestScore = score
			}
		}
	}

	fmt.Printf("day8 part1: %d\n", bestScore)
}

func viewDown(grid [][]int, startY, startX int) int {
	height := grid[startY][startX]
	var view int

	for y := startY + 1; y < len(grid); y++ {
		view++
		if grid[y][startX] >= height {
			return view
		}
	}
	return view
}

func viewUp(grid [][]int, startY, startX int) int {
	height := grid[startY][startX]
	var view int

	for y := startY - 1; y >= 0; y-- {
		view++
		if grid[y][startX] >= height {
			return view
		}
	}
	return view
}

func viewRight(grid [][]int, startY, startX int) int {
	height := grid[startY][startX]
	var view int

	for x := startX + 1; x < len(grid); x++ {
		view++
		if grid[startY][x] >= height {
			return view
		}
	}
	return view
}

func viewLeft(grid [][]int, startY, startX int) int {
	height := grid[startY][startX]
	var view int

	for x := startX - 1; x >= 0; x-- {
		view++
		if grid[startY][x] >= height {
			return view
		}
	}
	return view
}
