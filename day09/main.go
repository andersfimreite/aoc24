package main

import (
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
	checksum := 0

	var blocks []int

	for i, b := range data[:len(data)-1] {
		n := int(b - 48)
		_ = n
		id := -1
		if i%2 == 0 {
			id = i / 2
		}
		for j := 0; j < n; j += 1 {
			blocks = append(blocks, id)
		}
	}

	start := 0
	end := len(blocks) - 1

	for start <= end {
		if blocks[start] == -1 {
			for blocks[end] == -1 {
				end -= 1
			}
			if end < start {
				break
			}
			blocks[start] = blocks[end]
			end -= 1
		}

		checksum += start * blocks[start]
		start += 1
	}

	return checksum
}

func part2(data []byte) int {
	checksum := 0

	var blocks []int

	for i, b := range data[:len(data)-1] {
		n := int(b - 48)
		_ = n
		id := -1
		if i%2 == 0 {
			id = i / 2
		}
		for j := 0; j < n; j += 1 {
			blocks = append(blocks, id)
		}
	}

	end := len(blocks) - 1

	for end >= 0 {
		if blocks[end] == -1 {
			end -= 1
			continue
		}

		id := blocks[end]
		size := 1
		for end-size >= 0 && blocks[end-size] == id {
			size += 1
		}

		i := 0
		for i <= end-size {
			if blocks[i] != -1 {
				i += 1
				continue
			}

			gap := 1
			for blocks[i+gap] == -1 {
				gap += 1
			}

			if size <= gap {
				for j := 0; j < size; j += 1 {
					blocks[i+j] = id
					blocks[end-j] = -1
				}
				break
			}

			i += gap
		}
		end -= size
	}

	for i, id := range blocks {
		if id == -1 {
			continue
		}
		checksum += i * id
	}

	return checksum
}
