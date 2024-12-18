package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"strconv"
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

func part1(data []byte) string {
	rows := bytes.Split(data, []byte{'\n'})

	regA, _ := strconv.Atoi(string(bytes.Split(rows[0], []byte{':', ' '})[1]))
	regB, _ := strconv.Atoi(string(bytes.Split(rows[1], []byte{':', ' '})[1]))
	regC, _ := strconv.Atoi(string(bytes.Split(rows[2], []byte{':', ' '})[1]))

	prog := []int{}
	progRaw := bytes.Split(rows[4], []byte{':', ' '})[1]
	for _, p := range progRaw {
		if p != ',' {
			prog = append(prog, int(p-48))
		}
	}

	c := Computer{regA, regB, regC, 0}
	out := c.run(prog)

	str := ""
	for i, o := range out {
		str += strconv.Itoa(o)
		if i < len(out)-1 {
			str += ","
		}
	}
	return str
}

func part2(data []byte) int {
	rows := bytes.Split(data, []byte{'\n'})

	var prog []int
	progRaw := bytes.Split(rows[4], []byte{':', ' '})[1]
	for _, p := range progRaw {
		if p != ',' {
			prog = append(prog, int(p-48))
		}
	}

	possible := []int{0}
	idx := len(prog) - 1
	for idx >= 0 {
		possibleNew := []int{}
		for _, p := range possible {
			for i := range 8 {
				c := Computer{p<<3 + i, 0, 0, 0}
				out := c.run(prog)
				if len(out) == len(prog[idx:]) {
					ok := true
					for j, p := range prog[idx:] {
						if p != out[j] {
							ok = false
							break
						}
					}
					if ok {
						possibleNew = append(possibleNew, p<<3+i)
					}
				}
			}
		}
		possible = possibleNew
		idx -= 1
	}

	return possible[0]
}

type Computer struct {
	a, b, c int
	ip      int
}

func (c Computer) combo(operand int) int {
	if 0 <= operand && operand <= 3 {
		return operand
	}

	if operand == 4 {
		return c.a
	}
	if operand == 5 {
		return c.b
	}
	if operand == 6 {
		return c.c
	}

	panic("invalid operand")
}

func (c *Computer) run(prog []int) []int {
	out := []int{}

	for c.ip < len(prog) {
		opcode := Opcode(prog[c.ip])
		operand := prog[c.ip+1]

		switch opcode {
		case Adv:
			c.a = c.a / int(math.Pow(2, float64(c.combo(operand))))
		case Bxl:
			c.b = c.b ^ operand
		case Bst:
			c.b = c.combo(operand) % 8
		case Jnz:
			if c.a != 0 {
				c.ip = operand - 2 // -2 because 2 is added at the end
			}
		case Bxc:
			c.b = c.b ^ c.c
		case Out:
			out = append(out, c.combo(operand)%8)
		case Bdv:
			c.b = c.a / int(math.Pow(2, float64(c.combo(operand))))
		case Cdv:
			c.c = c.a / int(math.Pow(2, float64(c.combo(operand))))
		}

		c.ip += 2
	}

	return out
}

type Opcode int

const (
	Adv Opcode = iota
	Bxl
	Bst
	Jnz
	Bxc
	Out
	Bdv
	Cdv
)
