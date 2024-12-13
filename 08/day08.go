package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	inputraw, err := os.ReadFile("task.txt")
	//inputraw, err := os.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := [][]byte{}
	length := 0
	input = append(input, []byte{})
	for _, char := range inputraw {
		if char == '\n' {
			length++
			input = append(input, []byte{})
		} else {
			input[length] = append(input[length], byte(char))
		}
	}

	cpy1 := deepcopy(input)
	findAntinodes(cpy1)

	cpy2 := deepcopy(input)
	findAntinodesExt(cpy2)
	//printField(cpy2)
}

func printField(input [][]byte) {
	for x := range input {
		for y := range input[x] {
			fmt.Printf("%v", string(input[x][y]))
		}
		fmt.Printf("\n")
	}
}
func deepcopy(input [][]byte) [][]byte {
	cpy1 := make([][]byte, len(input))
	for i, row := range input {
		cpy1[i] = make([]byte, len(row))
		copy(cpy1[i], row)
	}
	return cpy1
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
func findAntennas(input [][]byte) []antenna {
	var ants []antenna
	for x := range input {
		for y := range len(input[0]) {
			if input[x][y] != '.' {
				ants = append(ants, antenna{x, y, input[x][y]})
			}
		}
	}
	return ants
}

type antenna struct {
	x     int
	y     int
	atype byte
}
type vector struct {
	x int
	y int
}

func findAntinodes(input [][]byte) {
	defer timerHighRes("Day 8 Part 1 findAntinodes")()

	xdim := len(input)
	ydim := len(input[0])
	ants := findAntennas(input)
	count := 0

	for i, ant := range ants {
		for _, candidate := range ants[i+1:] {
			if ant.atype != candidate.atype {
				continue
			}
			dx := candidate.x - ant.x
			dy := candidate.y - ant.y

			antatde := vector{ant.x - dx, ant.y - dy}
			if antatde.x < 0 || antatde.y < 0 || antatde.x >= xdim || antatde.y >= ydim {
			} else {
				if input[antatde.x][antatde.y] != '#' {
					input[antatde.x][antatde.y] = '#'
					count += 1
				}
			}
			candatde := vector{candidate.x + dx, candidate.y + dy}
			if candatde.x < 0 || candatde.y < 0 || candatde.x >= xdim || candatde.y >= ydim {
			} else {
				if input[candatde.x][candatde.y] != '#' {
					input[candatde.x][candatde.y] = '#'
					count += 1
				}
			}
		}
	}
	fmt.Printf("Count: %v\n", count)

}

func findAntinodesExt(input [][]byte) {
	defer timerHighRes("Day 8 Part 2 findAntinodesExtended")()

	xdim := len(input)
	ydim := len(input[0])
	ants := findAntennas(input)
	count := len(ants)

	for i, ant := range ants {
		for _, candidate := range ants[i+1:] {
			if ant.atype != candidate.atype {
				continue
			}
			dx := candidate.x - ant.x
			dy := candidate.y - ant.y
			r := 1
			for {
				antatde := vector{ant.x - dx*r, ant.y - dy*r}
				if antatde.x < 0 || antatde.y < 0 || antatde.x >= xdim || antatde.y >= ydim {
					break
				} else {
					if input[antatde.x][antatde.y] == '.' {
						input[antatde.x][antatde.y] = '#'
						count += 1
					}
				}
				r++
			}
			r = 1
			for {
				candatde := vector{candidate.x + dx*r, candidate.y + dy*r}
				if candatde.x < 0 || candatde.y < 0 || candatde.x >= xdim || candatde.y >= ydim {
					break
				} else {
					if input[candatde.x][candatde.y] == '.' {
						input[candatde.x][candatde.y] = '#'
						count += 1
					}
				}
				r++
			}
		}
	}
	fmt.Printf("Count: %v\n", count)

}
