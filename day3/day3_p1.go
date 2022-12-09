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

	for fileScanner.Scan() {
		items := strings.Split(fileScanner.Text(), "")
		leftItems := items[:len(items)/2]
		rightItems := strings.Join(items[len(items)/2:], "")

		for _, item := range leftItems {
			if strings.Contains(rightItems, item) {
				if int([]rune(item)[0]) < 97 {
					prioSum += int([]rune(item)[0]) - 38
				} else {
					prioSum += int([]rune(item)[0]) - 96
				}
				break
			}
		}
	}
	fmt.Printf("day3 part1: %d\n", prioSum)
}
