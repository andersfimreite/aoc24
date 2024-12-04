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

	rows := bytes.IndexByte(data, '\n') + 1
	cols := bytes.Count(data, []byte{'\n'}) + 1

	for i, b := range data {
		if b != 'X' && b != 'S' {
			continue
		}

		x := i % rows
		y := i / cols

		// check right
		if x < rows-4 {
			if b == 'X' {
				if data[i+1] == 'M' {
					if data[i+2] == 'A' {
						if data[i+3] == 'S' {
							count += 1
						}
					}
				}
			}

			if b == 'S' {
				if data[i+1] == 'A' {
					if data[i+2] == 'M' {
						if data[i+3] == 'X' {
							count += 1
						}
					}
				}
			}
		}

		// check down
		if y < cols-4 {
			if b == 'X' {
				if data[(y+1)*cols+x] == 'M' {
					if data[(y+2)*cols+x] == 'A' {
						if data[(y+3)*cols+x] == 'S' {
							count += 1
						}
					}
				}
			}

			if b == 'S' {
				if data[(y+1)*cols+x] == 'A' {
					if data[(y+2)*cols+x] == 'M' {
						if data[(y+3)*cols+x] == 'X' {
							count += 1
						}
					}
				}
			}
		}

		// check diag down right
		if x < rows-4 && y < cols-4 {
			if b == 'X' {
				if data[(y+1)*cols+x+1] == 'M' {
					if data[(y+2)*cols+x+2] == 'A' {
						if data[(y+3)*cols+x+3] == 'S' {
							count += 1
						}
					}
				}
			}

			if b == 'S' {
				if data[(y+1)*cols+x+1] == 'A' {
					if data[(y+2)*cols+x+2] == 'M' {
						if data[(y+3)*cols+x+3] == 'X' {
							count += 1
						}
					}
				}
			}
		}

		// check diag down left
		if 2 < x && y < cols-4 {
			if b == 'X' {
				if data[(y+1)*cols+(x-1)] == 'M' {
					if data[(y+2)*cols+(x-2)] == 'A' {
						if data[(y+3)*cols+(x-3)] == 'S' {
							count += 1
						}
					}
				}
			}

			if b == 'S' {
				if data[(y+1)*cols+(x-1)] == 'A' {
					if data[(y+2)*cols+(x-2)] == 'M' {
						if data[(y+3)*cols+(x-3)] == 'X' {
							count += 1
						}
					}
				}
			}
		}
	}

	return count
}

func part2(data []byte) int {
	count := 0

	rows := bytes.IndexByte(data, '\n') + 1
	cols := bytes.Count(data, []byte{'\n'}) + 1

	for i, b := range data {
		if b != 'A' {
			continue
		}

		x := i % rows
		y := i / cols

		if 0 < x && x < rows-2 && 0 < y && y < cols-2 {
			ul := data[(y-1)*cols+(x-1)]
			lr := data[(y+1)*cols+(x+1)]

			if ul == 'M' {
				if lr != 'S' {
					continue
				}
			} else if ul == 'S' {
				if lr != 'M' {
					continue
				}
			} else {
				continue
			}

			ur := data[(y-1)*cols+(x+1)]
			ll := data[(y+1)*cols+(x-1)]

			if ur == 'M' {
				if ll != 'S' {
					continue
				}
			} else if ur == 'S' {
				if ll != 'M' {
					continue
				}
			} else {
				continue
			}

			count += 1
		}
	}

	return count
}
