package main

import (
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

const Size = 71

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
	corrupted := make(map[string]int)
	coords := strings.Split(string(data), "\n")
	for i := range 1024 {
		corrupted[coords[i]] = 1
	}

	end := dijkstra(corrupted)

	return end.dist
}

func part2(data []byte) string {
	corrupted := make(map[string]int)
	coords := strings.Split(string(data), "\n")

	for i := range 1024 {
		corrupted[coords[i]] = 1
	}

	end := dijkstra(corrupted)

	for _, c := range coords[1024:] {
		corrupted[c] = 1

		prev := end.prev
		for prev != nil {
			key := fmt.Sprintf("%d,%d", prev.x, prev.y)
			if key == c {
				break
			}
			prev = prev.prev
		}

		if prev == nil {
			continue
		}

		end = dijkstra(corrupted)
		if end == nil {
			return c
		}
	}

	return ""
}

type Node struct {
	x, y int
	dist int
	prev *Node
}

func dijkstra(corrupted map[string]int) *Node {
	unvisited := []*Node{}

	for y := range Size {
		for x := range Size {
			key := fmt.Sprintf("%d,%d", x, y)
			if _, corr := corrupted[key]; !corr {
				dist := math.MaxInt
				if x == 0 && y == 0 {
					dist = 0
				}
				unvisited = append(unvisited, &Node{x, y, dist, nil})
			}
		}
	}

	for len(unvisited) != 0 {
		slices.SortFunc(unvisited, func(a, b *Node) int {
			return a.dist - b.dist
		})

		u := unvisited[0]
		unvisited = unvisited[1:]

		if u.dist == math.MaxInt {
			continue
		}

		if u.x == Size-1 && u.y == Size-1 {
			return u
		}

		d := u.dist + 1
		{
			nx := u.x - 1
			ny := u.y
			for _, v := range unvisited {
				if v.x == nx && v.y == ny {
					if d < v.dist {
						v.dist = d
						v.prev = u
					}
					break
				}
			}
		}
		{
			nx := u.x + 1
			ny := u.y
			for _, v := range unvisited {
				if v.x == nx && v.y == ny {
					if d < v.dist {
						v.dist = d
						v.prev = u
					}
					break
				}
			}
		}
		{
			nx := u.x
			ny := u.y - 1
			for _, v := range unvisited {
				if v.x == nx && v.y == ny {
					if d < v.dist {
						v.dist = d
						v.prev = u
					}
					break
				}
			}
		}
		{
			nx := u.x
			ny := u.y + 1
			for _, v := range unvisited {
				if v.x == nx && v.y == ny {
					if d < v.dist {
						v.dist = d
						v.prev = u
					}
					break
				}
			}
		}
	}

	return nil
}
