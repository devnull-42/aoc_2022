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
	maxCalCounts := make([]int, 3)

	for fileScanner.Scan() {
		cal, err := strconv.Atoi(fileScanner.Text())
		currentCalCount += cal

		if err != nil {
			switch {
			case currentCalCount > maxCalCounts[0]:
				maxCalCounts[2] = maxCalCounts[1]
				maxCalCounts[1] = maxCalCounts[0]
				maxCalCounts[0] = currentCalCount
			case currentCalCount > maxCalCounts[1]:
				maxCalCounts[2] = maxCalCounts[1]
				maxCalCounts[1] = currentCalCount
			case currentCalCount > maxCalCounts[2]:
				maxCalCounts[2] = currentCalCount
			}
			currentCalCount = 0
		}
	}

	var total int
	for _, count := range maxCalCounts {
		total += count
	}
	fmt.Printf("day1 part2: %d\n", total)
}
