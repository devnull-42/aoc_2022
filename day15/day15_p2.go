package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	lines := getData()

	posLines := make([]int, 0)
	negLines := make([]int, 0)

	for _, line := range lines {
		posLines = append(posLines, line[0]-line[1]+line[2])
		posLines = append(posLines, line[0]-line[1]-line[2])
		negLines = append(negLines, line[0]+line[1]+line[2])
		negLines = append(negLines, line[0]+line[1]-line[2])
	}

	sort.IntSlice.Sort(posLines)
	sort.IntSlice.Sort(negLines)

	pos, neg := 0, 0

	for i := 0; i < len(posLines)-2; i++ {
		if posLines[i]+2 == posLines[i+1] {
			pos = posLines[i] + 1
		}
		if negLines[i]+2 == negLines[i+1] {
			neg = negLines[i] + 1
		}
	}

	x := (pos + neg) / 2
	y := (neg - pos) / 2

	fmt.Printf("day15 part2: %d\n", x*4000000+y)

}

// getData returns a slice of the x, y and manhattan distance
// of a signal and it's beacon
func getData() [][]int {
	lines := make([][]int, 0)
	file, _ := os.Open("input.txt")
	defer file.Close()

	fs := bufio.NewScanner(file)
	fs.Split(bufio.ScanLines)

	re := regexp.MustCompile(`Sensor at x=(-?\d+), y=(-?\d+): closest beacon is at x=(-?\d+), y=(-?\d+)`)
	for fs.Scan() {
		match := re.FindStringSubmatch(fs.Text())

		signalX, _ := strconv.Atoi(match[1])
		signalY, _ := strconv.Atoi(match[2])
		beaconX, _ := strconv.Atoi(match[3])
		beaconY, _ := strconv.Atoi(match[4])

		distance := int(math.Abs(float64(signalX-beaconX)) + math.Abs(float64(signalY-beaconY)))

		lines = append(lines, []int{signalX, signalY, distance})
	}
	return lines
}
