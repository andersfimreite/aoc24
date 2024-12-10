package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
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

	rows := bytes.Split(bytes.TrimSuffix(data, []byte{'\n'}), []byte{'\n'})
	for y, row := range rows {
		for x, cell := range row {
			if cell != '0' {
				continue
			}

			seen := []int{}
			walk1(x, y, rows, &seen)

			score += len(seen)
		}
	}

	return score
}

func walk1(x, y int, rows [][]byte, seen *[]int) {
	curr := rows[y][x]

	if curr == '9' {
		idx := y*len(rows) + x
		if !slices.Contains(*seen, idx) {
			*seen = append(*seen, idx)
		}
		return
	}

	if 0 < x && rows[y][x-1] == curr+1 {
		walk1(x-1, y, rows, seen)
	}

	if x < len(rows)-1 && rows[y][x+1] == curr+1 {
		walk1(x+1, y, rows, seen)
	}

	if 0 < y && rows[y-1][x] == curr+1 {
		walk1(x, y-1, rows, seen)
	}

	if y < len(rows)-1 && rows[y+1][x] == curr+1 {
		walk1(x, y+1, rows, seen)
	}
}

func part2(data []byte) int {
	rating := 0

	rows := bytes.Split(bytes.TrimSuffix(data, []byte{'\n'}), []byte{'\n'})
	for y, row := range rows {
		for x, cell := range row {
			if cell != '0' {
				continue
			}

			seen := []int{}
			walk2(x, y, rows, &seen)

			rating += len(seen)
		}
	}

	return rating
}

func walk2(x, y int, rows [][]byte, seen *[]int) {
	curr := rows[y][x]

	if curr == '9' {
		idx := y*len(rows) + x
		*seen = append(*seen, idx)
		return
	}

	if 0 < x && rows[y][x-1] == curr+1 {
		walk2(x-1, y, rows, seen)
	}

	if x < len(rows)-1 && rows[y][x+1] == curr+1 {
		walk2(x+1, y, rows, seen)
	}

	if 0 < y && rows[y-1][x] == curr+1 {
		walk2(x, y-1, rows, seen)
	}

	if y < len(rows)-1 && rows[y+1][x] == curr+1 {
		walk2(x, y+1, rows, seen)
	}
}
