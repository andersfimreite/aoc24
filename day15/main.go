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
	parts := bytes.Split(data, []byte{'\n', '\n'})
	gridOld := bytes.Split(parts[0], []byte{'\n'})
	inst := parts[1]

	var r Vec2
	grid := [][]byte{}
	for y, rowOld := range gridOld {
		row := []byte{}
		for x, cell := range rowOld {
			if cell == '@' {
				r = Vec2{x, y}
			}
			row = append(row, cell)
		}
		grid = append(grid, row)
	}

	for _, i := range inst {
		if i == '\n' {
			continue
		}

		var cell byte
		var dir Vec2
		if i == '^' {
			cell = grid[r.y-1][r.x]
			dir = Vec2{0, -1}
		} else if i == 'v' {
			cell = grid[r.y+1][r.x]
			dir = Vec2{0, 1}
		} else if i == '<' {
			cell = grid[r.y][r.x-1]
			dir = Vec2{-1, 0}
		} else if i == '>' {
			cell = grid[r.y][r.x+1]
			dir = Vec2{1, 0}
		}

		if cell == '#' {
			continue
		}

		if cell == 'O' {
			pos := r
			pos.add(dir)
			for grid[pos.y][pos.x] == 'O' {
				pos.add(dir)
			}

			if grid[pos.y][pos.x] == '#' {
				continue
			}

			grid[pos.y][pos.x] = 'O'
		}

		grid[r.y][r.x] = '.'
		r.add(dir)
		grid[r.y][r.x] = '@'
	}

	sum := 0
	for y, row := range grid {
		for x, cell := range row {
			if cell == 'O' {
				sum += 100*y + x
			}
		}
	}

	return sum
}

func part2(data []byte) int {
	parts := bytes.Split(data, []byte{'\n', '\n'})
	gridOld := bytes.Split(parts[0], []byte{'\n'})
	inst := parts[1]

	var r Vec2
	grid := [][]byte{}
	for y, rowOld := range gridOld {
		row := []byte{}
		for x, cell := range rowOld {
			if cell == '@' {
				r = Vec2{x * 2, y}
				row = append(row, '@', '.')
			} else if cell == 'O' {
				row = append(row, '[', ']')
			} else {
				row = append(row, cell, cell)
			}
		}
		grid = append(grid, row)
	}

	for _, i := range inst {
		if i == '\n' {
			continue
		}

		var cell byte
		var dir Vec2
		if i == '^' {
			cell = grid[r.y-1][r.x]
			dir = Vec2{0, -1}
		} else if i == 'v' {
			cell = grid[r.y+1][r.x]
			dir = Vec2{0, 1}
		} else if i == '<' {
			cell = grid[r.y][r.x-1]
			dir = Vec2{-1, 0}
		} else if i == '>' {
			cell = grid[r.y][r.x+1]
			dir = Vec2{1, 0}
		}

		if cell == '#' {
			continue
		}

		if cell == '[' || cell == ']' {
			if i == '<' || i == '>' {
				pos := r
				pos.add(dir)
				for grid[pos.y][pos.x] == '[' || grid[pos.y][pos.x] == ']' {
					pos.add(dir)
					pos.add(dir)
				}

				if grid[pos.y][pos.x] == '#' {
					continue
				}

				for pos.x != r.x {
					tmp := pos
					tmp.sub(dir)
					grid[pos.y][pos.x] = grid[pos.y][tmp.x]
					pos.sub(dir)
				}
			} else {
				pos := r
				pos.add(dir)

				bounds := make(map[int][]int)
				if cell == '[' {
					left := r.x
					right := r.x + 1
					bounds[r.y] = []int{left, right}
					bounds[pos.y] = []int{left, right}
				} else {
					left := r.x - 1
					right := r.x
					bounds[r.y] = []int{left, right}
					bounds[pos.y] = []int{left, right}
				}

				canPush := true
			out:
				for true {
					left := bounds[pos.y][0]
					right := bounds[pos.y][1]
					pos.add(dir)

					for x := left; x <= right; x += 1 {
						if grid[pos.y][x] == '#' {
							canPush = false
							break out
						}
					}

					done := true
					for x := left; x <= right; x += 1 {
						if grid[pos.y][x] != '.' {
							done = false
							break
						}
					}
					if done {
						break
					}

					if grid[pos.y][left] == ']' {
						left -= 1
					} else if grid[pos.y][left] == '.' {
						for grid[pos.y][left] == '.' {
							left += 1
						}
					}
					if grid[pos.y][right] == '[' {
						right += 1
					} else if grid[pos.y][right] == '.' {
						for grid[pos.y][right] == '.' {
							right -= 1
						}
					}

					bounds[pos.y] = []int{left, right}
				}

				if canPush {
					for true {
						tmp := pos
						tmp.sub(dir)
						left := bounds[tmp.y][0]
						right := bounds[tmp.y][1]
						if tmp.y == r.y {
							for x := left; x <= right; x += 1 {
								grid[pos.y][x] = '.'
							}
							break
						}
						for x := left; x <= right; x += 1 {
							grid[pos.y][x] = grid[tmp.y][x]
							grid[tmp.y][x] = '.'
						}
						pos.sub(dir)
					}
				} else {
					continue
				}
			}
		}

		grid[r.y][r.x] = '.'
		r.add(dir)
		grid[r.y][r.x] = '@'
	}

	sum := 0
	for y, row := range grid {
		for x, cell := range row {
			if cell == '[' {
				sum += 100*y + x
			}
		}
	}

	return sum
}

type Vec2 struct {
	x, y int
}

func (v *Vec2) add(u Vec2) {
	v.x += u.x
	v.y += u.y
}

func (v *Vec2) sub(u Vec2) {
	v.x -= u.x
	v.y -= u.y
}

func printGrid(grid [][]byte) {
	for _, row := range grid {
		for _, cell := range row {
			fmt.Printf("%c", cell)
		}
		fmt.Println()
	}
}
