package main

import (
    "os"
    "fmt"
    "strings"
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

func part1(data []byte) int {
    safeCount := 0

    lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
    for _, line := range(lines) {
        nums := strings.Split(line, " ")

        safe := true

        prev := -1
        sign := 0
        for _, num := range(nums) {
            curr, _ := strconv.Atoi(num)
            if prev != -1 {
                diff := curr - prev

                if diff == 0 {
                    safe = false
                    break
                }

                if diff < 0 {
                    if sign == 0 {
                        sign = -1
                    }

                    if sign != -1 {
                        safe = false
                        break
                    }

                    diff = diff * -1
                } else {
                    if sign == 0 {
                        sign = 1
                    }

                    if sign != 1 {
                        safe = false
                        break
                    }
                }

                if diff > 3 {
                    safe = false
                    break
                }
            }

            prev = curr
        }

        if safe {
            safeCount += 1
        }
    }

    return safeCount;
}

func part2(data []byte) int {
    safeCount := 0

    lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
    for _, line := range(lines) {
        nums := strings.Split(line, " ")

        for skipIdx := 0; skipIdx < len(nums); skipIdx += 1 {
            safe := true

            prev := -1
            sign := 0
            for idx, num := range(nums) {
                if idx == skipIdx {
                    continue
                }

                curr, _ := strconv.Atoi(num)
                if prev != -1 {
                    diff := curr - prev

                    if diff == 0 {
                        safe = false
                        break
                    }

                    if diff < 0 {
                        if sign == 0 {
                            sign = -1
                        }

                        if sign != -1 {
                            safe = false
                            break
                        }

                        diff = diff * -1
                    } else {
                        if sign == 0 {
                            sign = 1
                        }

                        if sign != 1 {
                            safe = false
                            break
                        }
                    }

                    if diff > 3 {
                        safe = false
                        break
                    }
                }

                prev = curr
            }

            if safe {
                safeCount += 1
                break
            }
        }

    }

    return safeCount;
}
