package main

import (
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
    var left []int
    var right []int

    number := 0
    for _, b := range(data) {
        digit := int(b - 48)
        if 0 <= digit && digit <= 9 {
            number *= 10
            number += digit
        } else {
            if number == 0 {
                continue
            }
            if len(left) == len(right) {
                left = append(left, number)
            } else {
                right = append(right, number)
            }
            number = 0
        }
    }

    slices.Sort(left)
    slices.Sort(right)

    totalDist := 0
    for i := 0; i < len(left); i += 1 {
        dist := left[i] - right[i]
        if dist < 0 {
            dist *= -1
        }
        totalDist += dist
    }

    return totalDist
}

func part2(data []byte) int {
    var left []int
    counts := make(map[int]int)

    isLeft := true
    number := 0
    for _, b := range(data) {
        digit := int(b - 48)
        if 0 <= digit && digit <= 9 {
            number *= 10
            number += digit
        } else {
            if number == 0 {
                continue
            }
            if isLeft {
                left = append(left, number)
                isLeft = false
            } else {
                cnt, ok := counts[number]
                if !ok {
                    counts[number] = 1
                } else {
                    counts[number] = cnt + 1
                }
                isLeft = true
            }
            number = 0
        }
    }

    simScore := 0
    for _, n := range(left) {
        cnt := counts[n]
        simScore += n * cnt
    }

    return simScore
}
