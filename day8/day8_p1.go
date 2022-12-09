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

	var count int
	count += len(grid)*2 + (len(grid[0])-2)*2

	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[y])-1; x++ {
			if checkDown(grid, y, x) || checkUp(grid, y, x) || checkLeft(grid, y, x) || checkRight(grid, y, x) {
				count++
			}
		}
	}

	fmt.Printf("day8 part1: %d\n", count)
}

func checkDown(grid [][]int, startY, startX int) bool {
	height := grid[startY][startX]

	for y := startY + 1; y < len(grid); y++ {
		if grid[y][startX] >= height {
			return false
		}
	}
	return true
}

func checkUp(grid [][]int, startY, startX int) bool {
	height := grid[startY][startX]

	for y := startY - 1; y >= 0; y-- {
		if grid[y][startX] >= height {
			return false
		}
	}
	return true
}

func checkRight(grid [][]int, startY, startX int) bool {
	height := grid[startY][startX]

	for x := startX + 1; x < len(grid); x++ {
		if grid[startY][x] >= height {
			return false
		}
	}
	return true
}

func checkLeft(grid [][]int, startY, startX int) bool {
	height := grid[startY][startX]

	for x := startX - 1; x >= 0; x-- {
		if grid[startY][x] >= height {
			return false
		}
	}
	return true
}
