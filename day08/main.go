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
	count := 0

	rows := bytes.Split(data, []byte{'\n'})
	size := len(rows) - 1

	ants := make(map[byte][][2]int)

	for y, row := range rows {
		for x, freq := range row {
			if freq != '.' {
				ants[freq] = append(ants[freq], [2]int{x, y})
			}
		}
	}

	nodes := make(map[int]int)

	for _, positions := range ants {
		for i, a := range positions {
			for _, b := range positions[i+1:] {
				xDiff := b[0] - a[0]
				yDiff := b[1] - a[1]

				{
					x := a[0] - xDiff
					y := a[1] - yDiff

					if 0 <= x && x < size && 0 <= y && y < size {
						if _, ok := nodes[y*size+x]; !ok {
							count += 1
						}

						nodes[y*size+x] = 1
					}
				}

				{
					x := b[0] + xDiff
					y := b[1] + yDiff

					if 0 <= x && x < size && 0 <= y && y < size {
						if _, ok := nodes[y*size+x]; !ok {
							count += 1
						}

						nodes[y*size+x] = 1
					}
				}
			}
		}
	}

	return count
}

func part2(data []byte) int {
	count := 0

	rows := bytes.Split(data, []byte{'\n'})
	size := len(rows) - 1

	ants := make(map[byte][][2]int)

	for y, row := range rows {
		for x, freq := range row {
			if freq != '.' {
				ants[freq] = append(ants[freq], [2]int{x, y})
			}
		}
	}

	nodes := make(map[int]int)

	for _, positions := range ants {
		for i, a := range positions {
			for _, b := range positions[i+1:] {
				xDiff := b[0] - a[0]
				yDiff := b[1] - a[1]

				i := 0
				for true {
					x := a[0] - (xDiff * i)
					y := a[1] - (yDiff * i)

					if 0 <= x && x < size && 0 <= y && y < size {
						if _, ok := nodes[y*size+x]; !ok {
							count += 1
						}

						nodes[y*size+x] = 1

						i += 1
					} else {
						break
					}
				}

				i = 0
				for true {
					x := b[0] + (xDiff * i)
					y := b[1] + (yDiff * i)

					if 0 <= x && x < size && 0 <= y && y < size {
						if _, ok := nodes[y*size+x]; !ok {
							count += 1
						}

						nodes[y*size+x] = 1

						i += 1
					} else {
						break
					}
				}
			}
		}
	}

	return count
}
