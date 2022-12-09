package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	var currentCalCount int
	var maxCalCount int

	for fileScanner.Scan() {
		cal, err := strconv.Atoi(fileScanner.Text())
		currentCalCount += cal

		if err != nil {
			if currentCalCount > maxCalCount {
				maxCalCount = currentCalCount
			}
			currentCalCount = 0
		}
	}

	fmt.Printf("day1 part1: %d\n", maxCalCount)
}
