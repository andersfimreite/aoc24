package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const Rows = 101
const Cols = 103

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
	var q1, q2, q3, q4 int

	secs := 100

	r, _ := regexp.Compile("p=(-?[0-9]+),(-?[0-9]+) v=(-?[0-9]+),(-?[0-9]+)")
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")

	for _, line := range lines {
		match := r.FindStringSubmatch(line)
		px, _ := strconv.Atoi(match[1])
		py, _ := strconv.Atoi(match[2])
		vx, _ := strconv.Atoi(match[3])
		vy, _ := strconv.Atoi(match[4])

		dx := vx * secs
		dy := vy * secs

		nx := (px + dx) % Rows
		ny := (py + dy) % Cols

		if nx < 0 {
			nx += Rows
		}

		if ny < 0 {
			ny += Cols
		}

		if nx < Rows/2 {
			if ny < Cols/2 {
				q1 += 1
			}
			if ny > Cols/2 {
				q3 += 1
			}
			continue
		}

		if nx > Rows/2 {
			if ny < Cols/2 {
				q2 += 1
			}
			if ny > Cols/2 {
				q4 += 1
			}
			continue
		}
	}

	return q1 * q2 * q3 * q4
}

func part2(data []byte) int {
	r, _ := regexp.Compile("p=(-?[0-9]+),(-?[0-9]+) v=(-?[0-9]+),(-?[0-9]+)")
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")

	robots := [][]int{}
	for _, line := range lines {
		match := r.FindStringSubmatch(line)
		px, _ := strconv.Atoi(match[1])
		py, _ := strconv.Atoi(match[2])
		vx, _ := strconv.Atoi(match[3])
		vy, _ := strconv.Atoi(match[4])
		robots = append(robots, []int{px, py, vx, vy})
	}

	i := 0
	for i < 10000 {
		grid := [Cols][Rows]int{}
		for _, r := range robots {
			px := r[0]
			py := r[1]
			vx := r[2]
			vy := r[3]

			dx := vx * i
			dy := vy * i

			nx := (px + dx) % Rows
			ny := (py + dy) % Cols

			if nx < 0 {
				nx += Rows
			}

			if ny < 0 {
				ny += Cols
			}

			grid[ny][nx] += 1
		}

		tree := false
	out:
		for y, row := range grid {
			for x, cell := range row {
				if cell != 0 {
					s := 4
					if s-1 < x && x < Rows-s && y < Cols-s {
						t := true
						for k := 0; k < s; k += 1 {
							if grid[y+k][x-k] == 0 || grid[y+k][x+k] == 0 {
								t = false
								break
							}
						}
						tree = t
						if tree {
							break out
						}
					}
				}
			}
		}

		if tree {
			// for _, row := range grid {
			// 	for _, cell := range row {
			// 		if cell == 0 {
			// 			fmt.Print(".")
			// 		} else {
			// 			fmt.Print("0")
			// 		}
			// 	}
			// 	fmt.Println()
			// }
			// fmt.Println()
			// fmt.Println()
			break
		}

		i += 1
	}

	return i
}
