package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type monkey struct {
	items       []int
	op          func(old int) int
	test        func(val int) bool
	inspections int
	trueMonkey  int
	falseMonkey int
}

func (m *monkey) addItem(i int) {
	m.items = append(m.items, i)
}

func (m *monkey) popItem() int {
	pop := m.items[0]
	m.items = m.items[1:]
	return pop
}

func (m *monkey) turn(monkeys []*monkey, lcm int) {
	for len(m.items) > 0 {
		newItem := m.op(m.popItem()) % lcm
		m.inspections++
		if m.test(newItem) {
			monkeys[m.trueMonkey].addItem(newItem)
		} else {
			monkeys[m.falseMonkey].addItem(newItem)
		}
	}
}

func buildOperation(op string, x string) func(old int) int {
	switch op {
	case "+":
		switch {
		case x == "old":
			return func(old int) int {
				return old + old
			}
		default:
			val, _ := strconv.Atoi(x)
			return func(old int) int {
				return old + val
			}
		}
	default:
		switch {
		case x == "old":
			return func(old int) int {
				return old * old
			}
		default:
			val, _ := strconv.Atoi(x)
			return func(old int) int {
				return old * val
			}
		}
	}
}

func buildTest(x int) func(val int) bool {
	return func(val int) bool { return val%x == 0 }
}

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	fs := bufio.NewScanner(file)
	fs.Split(bufio.ScanLines)

	monkeys := make([]*monkey, 0)
	lcm := 1
	var m *monkey

	for fs.Scan() {
		if strings.HasPrefix(fs.Text(), "Monkey") {
			m = new(monkey)
			continue
		}
		if strings.HasPrefix(fs.Text(), "  Starting items:") {
			itemsString := strings.TrimLeft(fs.Text(), "  Starting items: ")
			items := strings.Split(itemsString, ", ")
			for _, i := range items {
				itemVal, _ := strconv.Atoi(i)
				m.addItem(itemVal)
			}
			continue
		}
		if strings.HasPrefix(fs.Text(), "  Operation:") {
			op := strings.Split(fs.Text(), " ")
			m.op = buildOperation(op[len(op)-2], op[len(op)-1])
			continue
		}
		if strings.HasPrefix(fs.Text(), "  Test:") {
			t := strings.Split(fs.Text(), " ")
			val, _ := strconv.Atoi(t[len(t)-1])
			lcm *= val
			m.test = buildTest(val)
			continue
		}
		if strings.HasPrefix(fs.Text(), "    If true:") {
			t := strings.Split(fs.Text(), " ")
			val, _ := strconv.Atoi(t[len(t)-1])
			m.trueMonkey = val
			continue
		}
		if strings.HasPrefix(fs.Text(), "    If false:") {
			t := strings.Split(fs.Text(), " ")
			val, _ := strconv.Atoi(t[len(t)-1])
			m.falseMonkey = val
			monkeys = append(monkeys, m)
			continue
		}
	}

	topInspections := []int{0, 0}
	rounds := 10000
	for j := 0; j < rounds; j++ {
		for i := 0; i < len(monkeys); i++ {
			monkeys[i].turn(monkeys, lcm)
			if j == rounds-1 {
				switch {
				case monkeys[i].inspections > topInspections[0]:
					topInspections[1] = topInspections[0]
					topInspections[0] = monkeys[i].inspections
				case monkeys[i].inspections > topInspections[1]:
					topInspections[1] = monkeys[i].inspections
				}
			}
		}
	}

	fmt.Printf("day11 part2: %d\n", topInspections[0]*topInspections[1])

}
