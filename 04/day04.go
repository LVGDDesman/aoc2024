package main

import (
	//"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	f, err := os.ReadFile("task.txt")
	//f, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := strings.Split(string(f), "\n")

	findxmas(input)
	findmasX(input)
}

func timerHighRes(name string) func() {
	start := time.Now()
	return func() {
		elapsed := time.Since(start)
		// Print duration with nanosecond precision
		fmt.Printf("%s took %d ns (%.6f ms)\n",
			name,
			elapsed.Nanoseconds(),
			float64(elapsed.Nanoseconds())/1_000_000.0)
	}
}

func findmasX(input []string) {
	defer timerHighRes("Day 4 2 findmasX")()
	count := 0
	for x := range len(input) {
		for y := range len(input[x]) {
			char := string(input[x][y])
			if char == "A" {
				if y+1 < len(input[x]) && y-1 >= 0 && x+1 < len(input) && x-1 >= 0 {
					if string(input[x-1][y-1]) == "S" && string(input[x+1][y+1]) == "M" || string(input[x-1][y-1]) == "M" && string(input[x+1][y+1]) == "S" {
						if string(input[x+1][y-1]) == "S" && string(input[x-1][y+1]) == "M" || string(input[x+1][y-1]) == "M" && string(input[x-1][y+1]) == "S" {
							count += 1
						}
					}
				}
			}
		}
	}
	fmt.Printf("%v\n", count)
}

func findxmas(input []string) {
	defer timerHighRes("Day 4 1 findxmas")()
	count := 0
	for x := range len(input) {
		for y := range len(input[x]) {
			char := string(input[x][y])
			if char == "X" {
				if y+3 < len(input[x]) && string(input[x][y+1]) == "M" && string(input[x][y+2]) == "A" && string(input[x][y+3]) == "S" {
					count += 1
				}
				if y-3 >= 0 && string(input[x][y-1]) == "M" && string(input[x][y-2]) == "A" && string(input[x][y-3]) == "S" {
					count += 1
				}
				if x+3 < len(input) && string(input[x+1][y]) == "M" && string(input[x+2][y]) == "A" && string(input[x+3][y]) == "S" {
					count += 1
				}
				if x-3 >= 0 && string(input[x-1][y]) == "M" && string(input[x-2][y]) == "A" && string(input[x-3][y]) == "S" {
					count += 1
				}
				if x-3 >= 0 && y-3 >= 0 && string(input[x-1][y-1]) == "M" && string(input[x-2][y-2]) == "A" && string(input[x-3][y-3]) == "S" {
					count += 1
				}
				if x-3 >= 0 && y+3 < len(input[x]) && string(input[x-1][y+1]) == "M" && string(input[x-2][y+2]) == "A" && string(input[x-3][y+3]) == "S" {
					count += 1
				}
				if x+3 < len(input) && y-3 >= 0 && string(input[x+1][y-1]) == "M" && string(input[x+2][y-2]) == "A" && string(input[x+3][y-3]) == "S" {
					count += 1
				}
				if x+3 < len(input) && y+3 < len(input[x]) && string(input[x+1][y+1]) == "M" && string(input[x+2][y+2]) == "A" && string(input[x+3][y+3]) == "S" {
					count += 1
				}

			}
		}
	}
	fmt.Printf("%v\n", count)
}
