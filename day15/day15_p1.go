package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	targetRow := 2000000
	file, _ := os.Open("input.txt")
	defer file.Close()

	fs := bufio.NewScanner(file)
	fs.Split(bufio.ScanLines)

	marked := make(map[int]interface{})
	re := regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)
	for fs.Scan() {
		match := re.FindStringSubmatch(fs.Text())

		signalX, _ := strconv.Atoi(match[1])
		signalY, _ := strconv.Atoi(match[2])
		beaconX, _ := strconv.Atoi(match[3])
		beaconY, _ := strconv.Atoi(match[4])

		distance := int(math.Abs(float64(signalX-beaconX)) + math.Abs(float64(signalY-beaconY)))

		if !(signalY-distance <= targetRow && signalY+distance >= targetRow) {
			continue
		}

		rangeX := distance - int(math.Abs(float64(signalY-targetRow)))

		for i := signalX - rangeX; i < signalX+rangeX; i++ {
			marked[i] = struct{}{}
		}

	}

	fmt.Printf("day15 part1: %d\n", len(marked))

}
