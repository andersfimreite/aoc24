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
	total := 0

	next := map[byte]byte{
		0:   'm',
		'm': 'u',
		'u': 'l',
		'l': '(',
	}

	var state byte = 0
	num1 := 0
	num2 := 0

	for _, b := range data {
		if n, ok := next[state]; ok {
			if n == b {
				state = n
			} else {
				state = 0
			}
			continue
		}

		if state == '(' {
			num1 = 0
			if '1' <= b && b <= '9' {
				num1 = int(b - 48)
				state = '1'
			} else {
				state = 0
			}
			continue
		}

		if state == '1' || state == '2' {
			if '0' <= b && b <= '9' {
				num1 *= 10
				num1 += int(b - 48)
				state += 1
			} else if b == ',' {
				state = ','
			} else {
				state = 0
			}
			continue
		}

		if state == '3' {
			if b == ',' {
				state = ','
			} else {
				state = 0
			}
			continue
		}

		if state == ',' {
			num2 = 0
			if '1' <= b && b <= '9' {
				num2 = int(b - 48)
				state = '4'
			} else {
				state = 0
			}
			continue
		}

		if state == '4' || state == '5' {
			if '0' <= b && b <= '9' {
				num2 *= 10
				num2 += int(b - 48)
				state += 1
			} else if b == ')' {
				total += num1 * num2
				state = 0
			} else {
				state = 0
			}
			continue
		}

		if state == '6' {
			if b == ')' {
				total += num1 * num2
			}
			state = 0
			continue
		}
	}

	return total
}

func part2(data []byte) int {
	total := 0

	next := map[byte]byte{
		'm':  'u',
		'u':  'l',
		'l':  '(',
		'd':  'o',
		'n':  '\'',
		'\'': 't',
	}

	var state byte = 0
	enabled := true
	nextEnabled := false
	num1 := 0
	num2 := 0

	for _, b := range data {
		if state == 0 {
			if b == 'd' {
				state = 'd'
			}
			if enabled && b == 'm' {
				state = 'm'
			}
			continue
		}

		if n, ok := next[state]; ok {
			if n == b {
				state = n
			} else {
				state = 0
			}
			continue
		}

		if state == 'o' {
			if b == '(' {
				state = '['
				nextEnabled = true
			} else if b == 'n' {
				state = 'n'
				nextEnabled = false
			} else {
				state = 0
			}
			continue
		}

		if state == 't' {
			if b == '(' {
				state = '['
			} else {
				state = 0
			}
			continue
		}

		if state == '[' {
			if b == ')' {
				enabled = nextEnabled
			}
			state = 0
			continue
		}

		if state == '(' {
			num1 = 0
			if '1' <= b && b <= '9' {
				num1 = int(b - 48)
				state = '1'
			} else {
				state = 0
			}
			continue
		}

		if state == '1' || state == '2' {
			if '0' <= b && b <= '9' {
				num1 *= 10
				num1 += int(b - 48)
				state += 1
			} else if b == ',' {
				state = ','
			} else {
				state = 0
			}
			continue
		}

		if state == '3' {
			if b == ',' {
				state = ','
			} else {
				state = 0
			}
			continue
		}

		if state == ',' {
			num2 = 0
			if '1' <= b && b <= '9' {
				num2 = int(b - 48)
				state = '4'
			} else {
				state = 0
			}
			continue
		}

		if state == '4' || state == '5' {
			if '0' <= b && b <= '9' {
				num2 *= 10
				num2 += int(b - 48)
				state += 1
			} else if b == ')' {
				total += num1 * num2
				state = 0
			} else {
				state = 0
			}
			continue
		}

		if state == '6' {
			if b == ')' {
				total += num1 * num2
			}
			state = 0
			continue
		}
	}

	return total
}
