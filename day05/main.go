package main

import (
	"fmt"
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
	sum := 0

	parts := strings.Split(string(data), "\n\n")
	rawRules := parts[0]
	rawUpdates := strings.TrimSuffix(parts[1], "\n")

	parsedRules := make(map[string][]string)

	rules := strings.Split(rawRules, "\n")
	for _, rule := range rules {
		parts := strings.Split(rule, "|")
		b := parts[0]
		a := parts[1]
		parsedRules[a] = append(parsedRules[a], b)
	}

	updates := strings.Split(rawUpdates, "\n")
	for _, update := range updates {
		pages := strings.Split(update, ",")
		ok := true
	outer:
		for i, p := range pages {
			for _, q := range pages[i+1:] {
				for _, v := range parsedRules[p] {
					if q == v {
						ok = false
						break outer
					}
				}
			}
		}

		if ok {
			middle, _ := strconv.Atoi(pages[len(pages)/2])
			sum += middle
		}
	}

	return sum
}

func part2(data []byte) int {
	sum := 0

	parts := strings.Split(string(data), "\n\n")
	rawRules := parts[0]
	rawUpdates := strings.TrimSuffix(parts[1], "\n")

	parsedRules := make(map[string][]string)

	rules := strings.Split(rawRules, "\n")
	for _, rule := range rules {
		parts := strings.Split(rule, "|")
		b := parts[0]
		a := parts[1]
		parsedRules[a] = append(parsedRules[a], b)
	}

	updates := strings.Split(rawUpdates, "\n")
	for _, update := range updates {
		pages := strings.Split(update, ",")
		ok := true
	outer:
		for i, p := range pages {
			for _, q := range pages[i+1:] {
				for _, v := range parsedRules[p] {
					if q == v {
						ok = false
						break outer
					}
				}
			}
		}

		if !ok {
			i := 0
		outer2:
			for i < len(pages) {
				for j, q := range pages[i+1:] {
					for _, v := range parsedRules[pages[i]] {
						if q == v {
							tmp := pages[i]
							pages[i] = pages[j+i+1]
							pages[j+i+1] = tmp
							continue outer2
						}
					}
				}
				i += 1
			}

			middle, _ := strconv.Atoi(pages[len(pages)/2])
			sum += middle
		}
	}

	return sum
}
