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
	price := 0

	rows := bytes.Split(bytes.TrimSuffix(data, []byte{'\n'}), []byte{'\n'})
	m := newMap(rows)

	for y, row := range m.rows {
		for x := range row {
			if !m.hasSeen(x, y) {
				area, peri := m.calc(x, y)
				price += area * peri
			}
		}
	}

	return price
}

func part2(data []byte) int {
	price := 0

	rows := bytes.Split(bytes.TrimSuffix(data, []byte{'\n'}), []byte{'\n'})
	m := newMap(rows)

	for y, row := range m.rows {
		for x := range row {
			if !m.hasSeen(x, y) {
				area, side := m.calc2(x, y)
				price += area * side
			}
		}
	}

	return price
}

type Map struct {
	rows [][]byte
	size int
	seen map[int]int
}

func newMap(rows [][]byte) Map {
	return Map{rows, len(rows), make(map[int]int)}
}

func (m Map) at(x, y int) byte {
	if x < 0 || y < 0 || x >= m.size || y >= m.size {
		return '.'
	}
	return m.rows[y][x]
}

func (m *Map) see(x, y int) {
	m.seen[y*m.size+x] = 1
}

func (m Map) hasSeen(x, y int) bool {
	_, ok := m.seen[y*m.size+x]
	return ok
}

func (m Map) calc(x, y int) (int, int) {
	area := 0
	peri := 0

	typ := m.rows[y][x]
	q := [][]int{[]int{x, y}}
	for len(q) > 0 {
		x := q[0][0]
		y := q[0][1]

		q = q[1:]

		if t := m.at(x, y); t != typ {
			peri += 1
			continue
		}

		if m.hasSeen(x, y) {
			continue
		}

		m.see(x, y)
		area += 1
		q = append(q, []int{x - 1, y}, []int{x + 1, y}, []int{x, y - 1}, []int{x, y + 1})
	}

	return area, peri
}

func (m Map) calc2(x, y int) (int, int) {
	area := 0
	side := 0

	typ := m.rows[y][x]
	q := [][]int{[]int{x, y}}
	for len(q) > 0 {
		x := q[0][0]
		y := q[0][1]

		q = q[1:]

		if t := m.at(x, y); t != typ {
			continue
		}

		if m.hasSeen(x, y) {
			continue
		}

		m.see(x, y)
		area += 1
		q = append(q, []int{x - 1, y}, []int{x + 1, y}, []int{x, y - 1}, []int{x, y + 1})

		u := m.at(x, y-1)
		d := m.at(x, y+1)
		l := m.at(x-1, y)
		r := m.at(x+1, y)

		if u != typ && l != typ {
			side += 1
		}

		if u != typ && r != typ {
			side += 1
		}

		if d != typ && l != typ {
			side += 1
		}

		if d != typ && r != typ {
			side += 1
		}

		ul := m.at(x-1, y-1)
		if u == typ && l == typ && ul != typ {
			side += 1
		}

		ur := m.at(x+1, y-1)
		if u == typ && r == typ && ur != typ {
			side += 1
		}

		dl := m.at(x-1, y+1)
		if d == typ && l == typ && dl != typ {
			side += 1
		}

		dr := m.at(x+1, y+1)
		if d == typ && r == typ && dr != typ {
			side += 1
		}
	}

	return area, side
}
