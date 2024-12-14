package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	filename := "input.txt"

	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	res1 := part1(data)
	fmt.Print("Part 1: ")
	fmt.Println(res1)

	res2 := part2(data)
	fmt.Print("Part 2: ")
	fmt.Println(res2)
}

func part1(data []byte) int {
	total := 0

	machines := bytes.Split(data, []byte{'\n', '\n'})
	for _, machine := range machines {
		m := parseMachine(machine, false)

		costs := []int{}
		is := []int{}

		i := -1
		for true {
			i += 1
			d := m.px - m.ax*i

			if i > 100 || d < 0 {
				break
			}

			if d%m.bx != 0 {
				continue
			}

			j := d / m.bx
			if j < 0 {
				break
			}

			if j > 100 {
				continue
			}

			cost := 3*i + j
			costs = append(costs, cost)
			is = append(is, i)
		}

		k := len(costs) - 1
		if k < 0 {
			continue
		}

		costBest := 0
		for k >= 0 {
			cost := costs[k]
			i := is[k]
			j := cost - i*3
			if m.ay*i+m.by*j == m.py {
				if costBest == 0 || cost < costBest {
					costBest = cost
				}
			}
			k -= 1
		}

		total += costBest
	}

	return total
}

func part2(data []byte) int {
	total := 0

	machines := bytes.Split(data, []byte{'\n', '\n'})
	for _, machine := range machines {
		m := parseMachine(machine, true)

		// Solve this with Cramers rule (https://en.wikipedia.org/wiki/Cramer%27s_rule)
		// m.ax * i + m.bx * j = m.px
		// m.ay * i + m.by * j = m.py

		d := m.ax*m.by - m.bx*m.ay
		di := m.px*m.by - m.bx*m.py
		dj := m.ax*m.py - m.px*m.ay

		// If the division doesn't yield a whole number the buttons can never get to the prize
		if di%d != 0 || dj%d != 0 {
			continue
		}

		i := di / d
		j := dj / d

		total += i*3 + j
	}

	return total
}

type Machine struct {
	ax, ay, bx, by, px, py int
}

func parseMachine(machine []byte, part2 bool) Machine {
	var ax, ay, bx, by, px, py int

	rows := bytes.Split(bytes.TrimSuffix(machine, []byte{'\n'}), []byte{'\n'})
	for i, row := range rows {
		s := byte('X')
		x := 0
		y := 0
		for _, c := range row {
			if c == 'Y' {
				s = c
				continue
			}

			n := int(c - 48)
			if 0 <= n && n <= 9 {
				if s == 'X' {
					x *= 10
					x += n
				} else {
					y *= 10
					y += n
				}
			}
		}

		if i == 0 {
			ax = x
			ay = y
		} else if i == 1 {
			bx = x
			by = y
		} else {
			if part2 {
				x += 10000000000000
				y += 10000000000000
			}
			px = x
			py = y
		}
	}

	return Machine{ax, ay, bx, by, px, py}
}
