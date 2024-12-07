package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
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

	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		test, _ := strconv.Atoi(parts[0])
		numStr := strings.Split(parts[1], " ")

		var nums []int
		for _, str := range numStr {
			n, _ := strconv.Atoi(str)
			nums = append(nums, n)
		}

		ok := rec1(test, nums[0], nums[1:])
		if ok {
			total += test
		}
	}

	return total
}

func rec1(test, sum int, nums []int) bool {
	if len(nums) == 0 {
		return test == sum
	}

	ok := rec1(test, sum+nums[0], nums[1:])
	if ok {
		return true
	}

	ok = rec1(test, sum*nums[0], nums[1:])

	return ok
}

func part2(data []byte) int {
	total := 0

	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		test, _ := strconv.Atoi(parts[0])
		numStr := strings.Split(parts[1], " ")

		var nums []int
		for _, str := range numStr {
			n, _ := strconv.Atoi(str)
			nums = append(nums, n)
		}

		ok := rec2(test, nums[0], nums[1:])
		if ok {
			total += test
		}
	}

	return total
}

func rec2(test, sum int, nums []int) bool {
	if len(nums) == 0 {
		return test == sum
	}

	ok := rec2(test, sum+nums[0], nums[1:])
	if ok {
		return true
	}

	ok = rec2(test, sum*nums[0], nums[1:])
	if ok {
		return true
	}

	digits := len(strconv.Itoa(nums[0]))
	ok = rec2(test, sum*int(math.Pow(10, float64(digits)))+nums[0], nums[1:])

	return ok
}
