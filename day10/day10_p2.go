package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type actions struct {
	ins []string
}

func (a *actions) pop() string {
	result := a.ins[0]
	a.ins = a.ins[1:]
	return result
}

func newActions() *actions {
	a := new(actions)
	a.ins = make([]string, 0)
	return a
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	fs := bufio.NewScanner(file)
	fs.Split(bufio.ScanLines)

	cycle := 0
	register := 1
	ready := 1

	a := newActions()

	for fs.Scan() {
		a.ins = append(a.ins, fs.Text())
	}

	for len(a.ins) > 0 {
		pos := cycle % 40
		if pos == 0 {
			fmt.Printf("\n")
		}
		if pos >= register-1 && pos <= register+1 {
			fmt.Printf("#")
		} else {
			fmt.Printf(".")
		}

		cycle++

		if ready == 0 {
			ins := a.pop()
			if ins != "noop" {
				action := strings.Split(ins, " ")
				val, _ := strconv.Atoi(action[1])
				register += val
				ready += 2
			} else {
				ready++
			}
		}

		ready--
	}
	fmt.Printf("day10 part1: image above\n")
}
