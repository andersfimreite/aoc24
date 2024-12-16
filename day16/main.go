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
	score := 0

	grid := bytes.Split(bytes.TrimSuffix(data, []byte{'\n'}), []byte{'\n'})
	start := Tile{1, len(grid) - 2, East, 0, nil}
	visited := make(map[string]int)
	q := []Tile{start}

	for len(q) > 0 {
		t := q[0]
		q = q[1:]

		visited[t.key()] = t.score

		if grid[t.y][t.x] == 'E' {
			if t.score < score || score == 0 {
				score = t.score
			}
			continue
		}

		for _, v := range t.next() {
			if grid[v.y][v.x] != '#' {
				s, found := visited[v.key()]
				if !found || v.score < s {
					q = append(q, v)
				}
			}
		}
	}

	return score
}

func part2(data []byte) int {
	score := 0

	grid := bytes.Split(bytes.TrimSuffix(data, []byte{'\n'}), []byte{'\n'})
	start := Tile{1, len(grid) - 2, East, 0, nil}
	visited := make(map[string]int)
	counted := make(map[string]int)
	q := []Tile{start}

	for len(q) > 0 {
		t := q[0]
		q = q[1:]

		visited[t.key()] = t.score

		if grid[t.y][t.x] == 'E' {
			if t.score < score || score == 0 {
				counted = make(map[string]int)
				score = t.score
			}
			prev := t.prev
			for prev != nil {
				if _, found := counted[prev.key()]; !found {
					counted[prev.key()] = 1
				}
				prev = prev.prev
			}
			continue
		}

		for _, v := range t.next() {
			if grid[v.y][v.x] != '#' {
				s, found := visited[v.key()]
				if !found || v.score < s {
					q = append(q, v)
				}
			}
		}
	}

	count := 1
	for _ = range counted {
		count += 1
	}

	return count
}

type Dir int

const (
	North Dir = iota
	South
	West
	East
)

type Tile struct {
	x, y  int
	dir   Dir
	score int
	prev  *Tile
}

func (t Tile) key() string {
	return fmt.Sprintf("%d-%d", t.x, t.y)
}

func (t Tile) next() []Tile {
	switch t.dir {
	case North:
		return []Tile{
			Tile{t.x, t.y - 1, North, t.score + 1, &t},
			Tile{t.x - 1, t.y, West, t.score + 1001, &t},
			Tile{t.x + 1, t.y, East, t.score + 1001, &t},
		}
	case South:
		return []Tile{
			Tile{t.x, t.y + 1, South, t.score + 1, &t},
			Tile{t.x - 1, t.y, West, t.score + 1001, &t},
			Tile{t.x + 1, t.y, East, t.score + 1001, &t},
		}
	case West:
		return []Tile{
			Tile{t.x, t.y - 1, North, t.score + 1001, &t},
			Tile{t.x, t.y + 1, South, t.score + 1001, &t},
			Tile{t.x - 1, t.y, West, t.score + 1, &t},
		}
	case East:
		return []Tile{
			Tile{t.x, t.y - 1, North, t.score + 1001, &t},
			Tile{t.x, t.y + 1, South, t.score + 1001, &t},
			Tile{t.x + 1, t.y, East, t.score + 1, &t},
		}
	}
	panic("unreachable")
}
