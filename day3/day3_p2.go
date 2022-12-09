package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	file, _ := os.Open("input.txt")
	defer file.Close()

	var prioSum int

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	itemSlice := make([]string, 0)

	for fileScanner.Scan() {
		itemSlice = append(itemSlice, fileScanner.Text())

		if len(itemSlice) == 3 {
			for _, item := range itemSlice[0] {
				if strings.Contains(itemSlice[1], string(item)) && strings.Contains(itemSlice[2], string(item)) {
					if int(item) < 97 {
						prioSum += int(item) - 38
					} else {
						prioSum += int(item) - 96
					}
					itemSlice = make([]string, 0)
					break
				}
			}
		}
	}
	fmt.Printf("day3 part2: %d\n", prioSum)
}
